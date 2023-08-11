package mock

import (
	"context"

	"github.com/think-free/sportsnews/lib/datamodel"
	"github.com/think-free/sportsnews/lib/logging"
	"github.com/think-free/sportsnews/providers/htafc/internal/cliparams"
)

// This is a basic mock database implementation that does nothing.
type Mock struct {
}

func New(ctx context.Context, _ *cliparams.ClientParameters) *Mock {
	return &Mock{}
}

func (m *Mock) SaveNews(ctx context.Context, news []datamodel.ICNews) error {
	logging.L(ctx).Infof("saving %d news", len(news))
	return nil
}
