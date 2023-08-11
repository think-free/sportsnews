package upstream_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/think-free/sportsnews/lib/logging"
	"github.com/think-free/sportsnews/providers/htafc/internal/cliparams"
	"github.com/think-free/sportsnews/providers/htafc/internal/config"
	"github.com/think-free/sportsnews/providers/htafc/internal/upstream"
)

func TestUpstream(t *testing.T) {
	ctx := context.Background()

	c := config.New()
	cp := cliparams.New()

	logging.Init(cp.LogLevel)

	u := upstream.New(ctx, cp, c)

	// Getting the news
	news, err := u.GetNews(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, news.Items.Articles)
	require.Empty(t, news.Items.Articles[0].BodyText)

	// Getting the first news id
	id := news.Items.Articles[0].ID
	require.NotEmpty(t, id)

	// Retreiving the full news for this id
	new, err := u.GetNewsByID(ctx, id)
	require.NoError(t, err)
	require.Equal(t, id, new.Article.ID)
	require.NotEmpty(t, new.Article.BodyText)
}
