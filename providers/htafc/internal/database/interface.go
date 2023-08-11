package database

import (
	"context"

	"github.com/think-free/sportsnews/lib/datamodel"
	"github.com/think-free/sportsnews/lib/logging"
	"github.com/think-free/sportsnews/providers/htafc/internal/cliparams"
	"github.com/think-free/sportsnews/providers/htafc/internal/database/mock"
	"github.com/think-free/sportsnews/providers/htafc/internal/database/mongodb"
)

type Database interface {
	SaveNews(ctx context.Context, news []datamodel.ICNews) error
}

// We instantiate a database implementation depending on the environment variable 'ENVIRONMENT'
func New(ctx context.Context, cp *cliparams.ClientParameters) Database {
	logging.L(ctx).Infof("creating database instance for environment '%s'", cp.Environment)

	switch cp.Environment {
	case "mock":
		if cp.MockedDatabase {
			return mock.New(ctx, cp)
		}
	case "production":
		return mongodb.New(ctx, cp)
	}
	return mongodb.New(ctx, cp)
}
