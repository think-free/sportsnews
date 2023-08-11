package mock

import (
	"context"
	"errors"

	"github.com/think-free/sportsnews/lib/datamodel"
	"github.com/think-free/sportsnews/lib/logging"
	"github.com/think-free/sportsnews/sportsnews/internal/cliparams"
	"github.com/think-free/sportsnews/sportsnews/internal/config"
)

// This is a basic mock database implementation for testing purpose, it should be improved.
type Mock struct {
	c *config.Config
}

func New(ctx context.Context, cp *cliparams.ClientParameters, c *config.Config) *Mock {
	return &Mock{
		c: c,
	}
}

func (m *Mock) GetNews(ctx context.Context, team string, page int) ([]*datamodel.ICNews, error) {
	logging.L(ctx).Infof("getting news from database for page '%d'", page)
	if page < 0 {
		return allArticles, nil
	}
	min := page * int(m.c.PageSize)
	max := min + int(m.c.PageSize)
	if max > len(allArticles) {
		max = len(allArticles)
	}
	return allArticles[min:max], nil
}

func (m *Mock) GetNewsByID(ctx context.Context, team string, id string) (*datamodel.ICNews, error) {
	logging.L(ctx).Infof("getting article '%s' from database", id)
	for i := range allArticles {
		if allArticles[i].ID == id {
			return allArticles[i], nil
		}
	}
	return nil, errors.New("article not found")
}
