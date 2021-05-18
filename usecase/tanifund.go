package usecase

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

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
	fundraisingID      = 6 // ID for "Fundraising"
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
	printer         *message.Printer
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
		printer:         message.NewPrinter(language.Indonesian),
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
		if notStarted(project.ProjectStatus.ID) && !tpc.projectCache[project.ID] {
			res := tpc.beautifyProject(project)
			if err := tpc.notifier.Notify(context.Background(), tpc.recipientID, res); err != nil {
				return err
			}
			tpc.projectCache[project.ID] = true
		}
	}
	return nil
}

func notStarted(id int) bool {
	return id == waitingForFundID || id == fundraisingID
}

func (tpc *TaniFundProjectChecker) beautifyProject(project *entity.Project) *entity.Project {
	project.HumanPublishedAt = project.PublishedAt.Add(TimeUTCPlus7).Format(time.RFC850)
	project.HumanCutoffAt = project.CutoffAt.Add(TimeUTCPlus7).Format(time.RFC850)
	project.ProjectLink = fmt.Sprintf("%s/%s", taniFundProjectURL, project.URLSlug)
	project.TargetFund = tpc.printer.Sprint(project.PricePerUnit * project.MaxUnit)
	project.Tenor = int(project.EndAt.Sub(project.StartAt) / thirtyDaysTime)
	return project
}
