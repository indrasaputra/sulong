package usecase

import (
	"context"

	"github.com/indrasaputra/sulong/entity"
)

const (
	// DefaultNumberOfProject set to be 10 because it is not too many, not too few.
	DefaultNumberOfProject = 10
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
	Notify(ctx context.Context, project *entity.Project) error
}

// TaniFundProjectChecker is responsible to check whether there is worthy new project to fund.
// If there is worthy new project to fund, it notifies the user.
type TaniFundProjectChecker struct {
	projectGetter   TaniFundProjectGetter
	notifier        TaniFundProjectNotifier
	numberOfProject int
	projectCache    map[string]bool
}

// NewTaniFundProjectChecker creates an instance of TaniFundProjectChecker.
// By default, the number of project to be retrieved is DefaultNumberOfProject.
func NewTaniFundProjectChecker(projectGetter TaniFundProjectGetter, notifier TaniFundProjectNotifier) *TaniFundProjectChecker {
	return &TaniFundProjectChecker{
		projectGetter:   projectGetter,
		notifier:        notifier,
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
		if !tpc.projectCache[project.ID] {
			if err := tpc.notifier.Notify(context.Background(), project); err != nil {
				return err
			}
			tpc.projectCache[project.ID] = true
		}
	}
	return nil
}
