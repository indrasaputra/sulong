package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/sulong/entity"
	mock_usecase "github.com/indrasaputra/sulong/test/mock/usecase"
	"github.com/indrasaputra/sulong/usecase"
)

var (
	ctx         = context.Background()
	errCustom   = errors.New("error")
	recipientID = 1
	projects    = []*entity.Project{
		{ID: "1", ProjectStatus: &entity.ProjectStatus{ID: 5}},
		{ID: "2", ProjectStatus: &entity.ProjectStatus{ID: 6}},
		{ID: "3", ProjectStatus: &entity.ProjectStatus{ID: 3}},
	}
)

type TaniFundProjectCheckerExecutor struct {
	checker  *usecase.TaniFundProjectChecker
	getter   *mock_usecase.MockTaniFundProjectGetter
	notifier *mock_usecase.MockTaniFundProjectNotifier
}

func TestNewTaniFundProjectChecker(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of TaniFundProjectChecker", func(t *testing.T) {
		exec := createTaniFundProjectCheckerExecutor(ctrl)
		assert.NotNil(t, exec.checker)
	})
}

func TestTaniFundProjectChecker_CheckAndNotify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("projects can't be retrieved", func(t *testing.T) {
		exec := createTaniFundProjectCheckerExecutor(ctrl)
		exec.getter.EXPECT().GetNewestProjects(ctx, usecase.DefaultNumberOfProject).Return([]*entity.Project{}, errCustom)

		err := exec.checker.CheckAndNotify()

		assert.NotNil(t, err)
	})

	t.Run("notifier returns error", func(t *testing.T) {
		exec := createTaniFundProjectCheckerExecutor(ctrl)
		exec.getter.EXPECT().GetNewestProjects(ctx, usecase.DefaultNumberOfProject).Return(projects, nil)
		exec.notifier.EXPECT().Notify(ctx, recipientID, projects[0]).Return(errCustom)

		err := exec.checker.CheckAndNotify()

		assert.NotNil(t, err)
	})

	t.Run("successfully checks and notify", func(t *testing.T) {
		exec := createTaniFundProjectCheckerExecutor(ctrl)
		exec.getter.EXPECT().GetNewestProjects(ctx, usecase.DefaultNumberOfProject).Return(projects, nil)
		exec.notifier.EXPECT().Notify(ctx, recipientID, projects[0]).Return(nil)
		exec.notifier.EXPECT().Notify(ctx, recipientID, projects[1]).Return(nil)

		err := exec.checker.CheckAndNotify()

		assert.Nil(t, err)
	})
}

func TestTaniFundProjectChecker_SetNumberOfProject(t *testing.T) {
	t.Run("successfully sets the number of project", func(t *testing.T) {
		checker := usecase.NewTaniFundProjectChecker(nil, nil, recipientID)

		assert.NotPanics(t, func() { checker.SetNumberOfProject(1) })
	})
}

func createTaniFundProjectCheckerExecutor(ctrl *gomock.Controller) *TaniFundProjectCheckerExecutor {
	getter := mock_usecase.NewMockTaniFundProjectGetter(ctrl)
	notifier := mock_usecase.NewMockTaniFundProjectNotifier(ctrl)

	return &TaniFundProjectCheckerExecutor{
		getter:   getter,
		notifier: notifier,
		checker:  usecase.NewTaniFundProjectChecker(getter, notifier, recipientID),
	}
}
