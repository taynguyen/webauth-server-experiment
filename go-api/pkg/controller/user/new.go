package user

import (
	"context"

	"github.com/dwarvesf/go-api/pkg/config"
	"github.com/dwarvesf/go-api/pkg/logger/monitor"
	"github.com/dwarvesf/go-api/pkg/model"
	"github.com/dwarvesf/go-api/pkg/repository"
)

// Controller auth controller
type Controller interface {
	Me(ctx context.Context) (*model.User, error)
	UpdateUser(ctx context.Context, user model.UpdateUserRequest) (*model.User, error)
	UpdatePassword(ctx context.Context, user model.UpdatePasswordRequest) error
	SentMail(ctx context.Context) error
}

type impl struct {
	repo    *repository.Repo
	cfg     config.Config
	monitor monitor.Tracer
}

// NewUserController new auth controller
func NewUserController(cfg config.Config, r *repository.Repo, monitor monitor.Tracer) Controller {
	return &impl{
		repo:    r,
		cfg:     cfg,
		monitor: monitor,
	}
}
