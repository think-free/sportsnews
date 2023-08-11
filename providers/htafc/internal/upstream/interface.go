package upstream

import (
	"context"

	"github.com/think-free/sportsnews/lib/logging"
	"github.com/think-free/sportsnews/providers/htafc/internal/cliparams"
	"github.com/think-free/sportsnews/providers/htafc/internal/config"
	"github.com/think-free/sportsnews/providers/htafc/internal/upstream/client"
	"github.com/think-free/sportsnews/providers/htafc/internal/upstream/mock"
	"github.com/think-free/sportsnews/providers/htafc/internal/upstream/model"
)

// Upstream is the interface that retrieve the news from 'Huddersfield Town' api
// The implementation can be the real client, a mock client, or a simulated client for testing purposes
type Upstream interface {
	GetNews(ctx context.Context) (*model.NewsArticles, error)               // Get all the news
	GetNewsByID(ctx context.Context, id string) (*model.NewsArticle, error) // Get a specific news with more details
}

// We instantiate an upstream implementation depending on the environment variable 'ENVIROMENT'
func New(ctx context.Context, cp *cliparams.ClientParameters, c *config.Config) Upstream {
	logging.L(ctx).Infof("creating upstream instance for environment '%s'", cp.Environment)

	switch cp.Environment {
	case "mock":
		if cp.MockedUpstream {
			return mock.New(ctx, c)
		}
	case "production":
		return client.New(ctx, c)
	}
	return client.New(ctx, c)
}
