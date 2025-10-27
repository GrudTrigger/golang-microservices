package app

import (
	"context"

	"github.com/go-faster/errors"
	"github.com/rocker-crm/notifacation/internal/config"
	"github.com/rocker-crm/platform/pkg/closer"
	"github.com/rocker-crm/platform/pkg/logger"
	"go.uber.org/zap"
)

type App struct {
	diContainer *diContainer
}

func New(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	errCh := make(chan error, 2)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		if err := a.runConsumerOrderPaid(ctx); err != nil {
			errCh <- errors.Errorf("consumer crashed: %v", err)
		}
	}()

	go func() {
		if err := a.runConsumerShipAssembled(ctx); err != nil {
			errCh <- errors.Errorf("consumer crashed: %v", err)
		}
	}()

	select {
	case <-ctx.Done():
		logger.Info(ctx, "Shutdown signal received")
	case err := <-errCh:
		logger.Error(ctx, "Component crashed, shutting down", zap.Error(err))
		cancel()
		<-ctx.Done()
		return err
	}
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initDI,
		a.initLogger,
		a.initCloser,
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
	return logger.Init(
		config.AppConfig().Logger.Level(),
		config.AppConfig().Logger.AsJson(),
	)
}

func (a *App) initCloser(_ context.Context) error {
	closer.SetLogger(logger.Logger())
	return nil
}

func (a *App) runConsumerOrderPaid(ctx context.Context) error {
	logger.Info(ctx, "🚀 Order-Paid Kafka consumer running")

	err := a.diContainer.OrderPaidConsumerService().RunConsumer(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runConsumerShipAssembled(ctx context.Context) error {
	logger.Info(ctx, "🚀 Ship-Assembled Kafka consumer running")

	err := a.diContainer.ShipAssembledConsumerService().RunConsumer(ctx)
	if err != nil {
		return err
	}

	return nil
}
