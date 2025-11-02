package app

import (
	"context"
	"fmt"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/rocker-crm/platform/pkg/cache"
	"github.com/rocker-crm/platform/pkg/cache/redis"
	"github.com/rocker-crm/platform/pkg/closer"
	"github.com/rocker-crm/platform/pkg/logger"
	authV1 "github.com/rocker-crm/shared/pkg/proto/auth/v1"
	authAPIV1 "github.com/rocket-crm/iam/internal/api/auth/v1"
	"github.com/rocket-crm/iam/internal/config"
	"github.com/rocket-crm/iam/internal/migrator"
	"github.com/rocket-crm/iam/internal/repository"
	"github.com/rocket-crm/iam/internal/repository/session"
	"github.com/rocket-crm/iam/internal/repository/user"
	"github.com/rocket-crm/iam/internal/service"
	"github.com/rocket-crm/iam/internal/service/auth"
)

type diContainer struct {
	authV1API         authV1.AuthServiceServer
	authService       service.AuthService
	authRepository    repository.UserRepository
	sessionRepository repository.SessionRepository

	postgresDb  *pgx.Conn
	redisPool   *redigo.Pool
	redisClient cache.RedisClient
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) AuthV1API(ctx context.Context) authV1.AuthServiceServer {
	if d.authV1API == nil {
		d.authV1API = authAPIV1.NewApi(d.AuthService(ctx))
	}
	return d.authV1API
}

func (d *diContainer) AuthService(ctx context.Context) service.AuthService {
	if d.authService == nil {
		d.authService = auth.NewService(d.AuthRepository(ctx), d.SessionRepository())
	}
	return d.authService
}

func (d *diContainer) AuthRepository(ctx context.Context) repository.UserRepository {
	if d.authRepository == nil {
		d.authRepository = user.NewRepository(d.PostgresDb(ctx))
	}
	return d.authRepository
}

func (d *diContainer) SessionRepository() repository.SessionRepository {
	if d.sessionRepository == nil {
		d.sessionRepository = session.NewRepository(d.RedisClient())
	}
	return d.sessionRepository
}

func (d *diContainer) PostgresDb(ctx context.Context) *pgx.Conn {
	if d.postgresDb == nil {
		conn, err := pgx.Connect(ctx, config.AppConfig().Postgres.URI())
		if err != nil {
			panic(fmt.Errorf("failed to connect to database: %w", err))
		}

		err = conn.Ping(ctx)
		if err != nil {
			panic(fmt.Errorf("failed to connect to database: %w", err))
		}
		migratorRunner := migrator.NewMigrator(stdlib.OpenDB(*conn.Config().Copy()), config.AppConfig().Postgres.MigrationDir())

		err = migratorRunner.Up()
		if err != nil {
			panic(fmt.Errorf("failed migrations up: %w", err))
		}

		closer.AddNamed("Postgres connect", func(ctx context.Context) error {
			return conn.Close(ctx)
		})

		d.postgresDb = conn
	}
	return d.postgresDb
}

func (d *diContainer) RedisPool() *redigo.Pool {
	if d.redisPool == nil {
		d.redisPool = &redigo.Pool{
			MaxIdle:     config.AppConfig().Redis.MaxIdle(),
			IdleTimeout: config.AppConfig().Redis.IdleTimeout(),
			DialContext: func(ctx context.Context) (redigo.Conn, error) {
				return redigo.DialContext(ctx, "tcp", config.AppConfig().Redis.Address())
			},
		}
	}
	return d.redisPool
}

func (d *diContainer) RedisClient() cache.RedisClient {
	if d.redisClient == nil {
		d.redisClient = redis.NewClient(d.RedisPool(), logger.Logger(), config.AppConfig().Redis.ConnectionTimeout())
	}
	return d.redisClient
}
