package pwl

import (
	"github.com/dwarvesf/go-api/pkg/config"
	"github.com/dwarvesf/go-api/pkg/controller/auth"
	"github.com/dwarvesf/go-api/pkg/controller/user"
	"github.com/dwarvesf/go-api/pkg/logger"
	"github.com/dwarvesf/go-api/pkg/logger/monitor"
	"github.com/dwarvesf/go-api/pkg/repository"
	"github.com/dwarvesf/go-api/pkg/service"
	"github.com/go-webauthn/webauthn/webauthn"
)

// Handler for app
type Handler struct {
	cfg      config.Config
	log      logger.Log
	svc      service.Service
	monitor  monitor.Tracer
	authCtrl auth.Controller
	userCtrl user.Controller
	webAuthn *webauthn.WebAuthn
}

// New will return an instance of Auth struct
func New(
	cfg config.Config,
	l logger.Log,
	repo *repository.Repo,
	svc service.Service,
	webAuthn *webauthn.WebAuthn,
	monitor monitor.Tracer) *Handler {
	return &Handler{
		cfg:      cfg,
		log:      l,
		svc:      svc,
		monitor:  monitor,
		authCtrl: auth.NewAuthController(cfg, repo, monitor),
		userCtrl: user.NewUserController(cfg, repo, monitor),
		webAuthn: webAuthn,
	}
}
