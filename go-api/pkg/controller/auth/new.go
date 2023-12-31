package auth

import (
	"context"

	"github.com/dwarvesf/go-api/pkg/config"
	"github.com/dwarvesf/go-api/pkg/logger/monitor"
	"github.com/dwarvesf/go-api/pkg/model"
	"github.com/dwarvesf/go-api/pkg/repository"
	"github.com/dwarvesf/go-api/pkg/service/jwthelper"
	"github.com/dwarvesf/go-api/pkg/service/passwordhelper"
)

// Controller auth controller
type Controller interface {
	Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error)
	Signup(ctx context.Context, req model.SignupRequest) error
}

type impl struct {
	repo           *repository.Repo
	jwtHelper      jwthelper.Helper
	cfg            config.Config
	monitor        monitor.Tracer
	passwordHelper passwordhelper.Helper
}

// NewAuthController new auth controller
func NewAuthController(cfg config.Config, r *repository.Repo, monitor monitor.Tracer) Controller {
	return &impl{
		repo:           r,
		jwtHelper:      jwthelper.NewHelper(cfg.SecretKey),
		cfg:            cfg,
		monitor:        monitor,
		passwordHelper: passwordhelper.NewScrypt(),
	}
}
