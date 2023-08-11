package database

import (
	"context"

	"github.com/think-free/sportsnews/lib/datamodel"
	"github.com/think-free/sportsnews/lib/logging"
	"github.com/think-free/sportsnews/sportsnews/internal/cliparams"
	"github.com/think-free/sportsnews/sportsnews/internal/config"
	"github.com/think-free/sportsnews/sportsnews/internal/database/mock"
	"github.com/think-free/sportsnews/sportsnews/internal/database/mongodb"
)

type Database interface {
	// GetNews returns a page of news for a given team, if team is not specified, it returns news for all teams, if page is not specified, it returns all the news
	GetNews(ctx context.Context, team string, page int) ([]*datamodel.ICNews, error)
	// GetNewsByID returns a news for a given team and id, team can be omitted, it will return the news with the given id
	GetNewsByID(ctx context.Context, team string, id string) (*datamodel.ICNews, error)
}

func New(ctx context.Context, cp *cliparams.ClientParameters, c *config.Config) Database {
	logging.L(ctx).Infof("creating database instance for environment '%s'", cp.Environment)

	switch cp.Environment {
	case "production":
		return mongodb.New(ctx, cp, c)
	case "mock":
		return mock.New(ctx, cp, c)
	}

	return mongodb.New(ctx, cp, c)
}
