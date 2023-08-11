package mock

import (
	"context"

	"github.com/think-free/sportsnews/providers/htafc/internal/config"
	"github.com/think-free/sportsnews/providers/htafc/internal/upstream/model"
)

// Mock is a mock implementation of the upstream interface
// It is used for testing purposes
// It implements the Upstream interface and can be used by the service
// It returns a predefined set of articles

type Mock struct {
}

func New(_ context.Context, _ *config.Config) *Mock {
	return &Mock{}
}

func (m *Mock) GetNews(_ context.Context) (*model.NewsArticles, error) {
	return articles, nil
}

func (m *Mock) GetNewsByID(_ context.Context, id string) (*model.NewsArticle, error) {
	return article, nil
}
