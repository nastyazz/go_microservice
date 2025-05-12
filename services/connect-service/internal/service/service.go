package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nastyazz/go_microservice.git/internal/proxyproto"
	sqlc "github.com/nastyazz/go_microservice.git/internal/userdb"
	"github.com/nastyazz/go_microservice.git/services/connect-service/internal/config"
	"github.com/nastyazz/go_microservice.git/services/connect-service/internal/kc"
)

type Service struct {
	proxyproto.UnimplementedCentrifugoProxyServer
	conn     *pgxpool.Pool
	queries  *sqlc.Queries
	kcClient *kc.KCClient
}

func New(cfg *config.Config) (*Service, error) {
	connCfg, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), connCfg)
	if err != nil {
		return nil, err
	}
	kcClient := kc.New(cfg.KeyCloakURL, cfg.KeyCloakRealm, cfg.KeyCloakClient, cfg.KeyCloakSecret)
	return &Service{
		conn:     conn,
		queries:  sqlc.New(conn),
		kcClient: kcClient,
	}, nil
}
