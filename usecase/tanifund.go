package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/indrasaputra/sulong/entity"
)

const (
	// DefaultNumberOfProject set to be 10 because it is not too many, not too few.
	DefaultNumberOfProject = 10
	// TimeUTCPlus7 derived from UTC+7 (Asia/Jakarta).
	TimeUTCPlus7 = 7 * time.Hour

	taniFundProjectURL = "https://tanifund.com/project"
	thirtyDaysTime     = 30 * 24 * time.Hour
	waitingForFundID   = 5 // ID for "Menunggu Fundraising"
)

// TaniFundProjectGetter defines a contract to get TaniFund's projects.
type TaniFundProjectGetter interface {
	// GetNewestProjects gets n newest TaniFund's projects.
	// The number of projcet can be set via parameter.
	GetNewestProjects(ctx context.Context, numberOfProject int) ([]*entity.Project, error)
}

// TaniFundProjectNotifier defines a contract to user about TaniFund's project.
type TaniFundProjectNotifier interface {
	// Notify notifies user about a project.
	Notify(ctx context.Context, recipientID int, project *entity.Project) error
}

// TaniFundProjectChecker is responsible to check whether there is worthy new project to fund.
// If there is worthy new project to fund, it notifies the user.
type TaniFundProjectChecker struct {
	projectGetter   TaniFundProjectGetter
	notifier        TaniFundProjectNotifier
	recipientID     int
	numberOfProject int
	projectCache    map[string]bool
}

// NewTaniFundProjectChecker creates an instance of TaniFundProjectChecker.
// By default, the number of project to be retrieved is DefaultNumberOfProject.
func NewTaniFundProjectChecker(projectGetter TaniFundProjectGetter, notifier TaniFundProjectNotifier, recipientID int) *TaniFundProjectChecker {
	return &TaniFundProjectChecker{
		projectGetter:   projectGetter,
		notifier:        notifier,
		recipientID:     recipientID,
		numberOfProject: DefaultNumberOfProject,
		projectCache:    make(map[string]bool),
	}
}

// SetNumberOfProject sets the number of project to be retrieved at once.
func (tpc *TaniFundProjectChecker) SetNumberOfProject(n int) {
	tpc.numberOfProject = n
}

// CheckAndNotify checks TaniFund's projects and notify if there is new project.
func (tpc *TaniFundProjectChecker) CheckAndNotify() error {
	projects, err := tpc.projectGetter.GetNewestProjects(context.Background(), tpc.numberOfProject)
	if err != nil {
		return err
	}

	for _, project := range projects {
		if project.Projectstatus.ID == waitingForFundID && !tpc.projectCache[project.ID] {
			res := beautifyProject(project)
			if err := tpc.notifier.Notify(context.Background(), tpc.recipientID, res); err != nil {
				return err
			}
			tpc.projectCache[project.ID] = true
		}
	}
	return nil
}

func beautifyProject(project *entity.Project) *entity.Project {
	project.HumanPublishedAt = project.PublishedAt.Add(TimeUTCPlus7)
	project.ProjectLink = fmt.Sprintf("%s/%s", taniFundProjectURL, project.URLSlug)
	project.TargetFund = project.PricePerUnit * project.MaxUnit
	project.Tenor = int(project.EndAt.Sub(project.StartAt) / thirtyDaysTime)
	return project
}
