package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-faster/errors"
	"github.com/rocker-crm/platform/pkg/closer"
	"github.com/rocker-crm/platform/pkg/logger"
	ordersV1 "github.com/rocker-crm/shared/pkg/openapi/orders/v1"
	"github.com/rocket-crm/order/internal/config"
)

type App struct {
	diContainer *diContainer
	httpServer  *ordersV1.Server
	server      *http.Server
}

func New(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initDI,
		a.initLogger,
		a.initCloser,
		a.initConfigServer,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initDI(_ context.Context) error {
	a.diContainer = NewDiContainer()
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	return logger.Init(config.AppConfig().Logger.Level(), config.AppConfig().Logger.AsJson())
}

func (a *App) initCloser(_ context.Context) error {
	closer.SetLogger(logger.Logger())
	return nil
}

func (a *App) initConfigServer(_ context.Context) error {
	r := chi.NewRouter()

	// Добавляем middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	// Монтируем обработчики OpenAPI
	r.Mount("/", a.httpServer)

	// Запускаем HTTP-сервер
	server := &http.Server{
		Addr:              config.AppConfig().OrderHttp.Address(),
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	a.server = server

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	ordersServer, err := ordersV1.NewServer(a.diContainer.OrderV1API(ctx))
	if err != nil {
		return err
	}
	a.httpServer = ordersServer
	return nil
}

func (a *App) runHTTPServer(ctx context.Context) error {
	logger.Info(ctx, fmt.Sprintf("🚀 HTTP-сервер запущен на порту %s\n", config.AppConfig().OrderHttp.Address()))
	err := a.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	closer.AddNamed("HTTP Server", func(ctx context.Context) error {
		return a.server.Shutdown(ctx)
	})
	return nil
}

func (a *App) Run(ctx context.Context) error {
	return a.runHTTPServer(ctx)
}
