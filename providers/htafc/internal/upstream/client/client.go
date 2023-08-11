package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/think-free/sportsnews/lib/restylog"
	"github.com/think-free/sportsnews/providers/htafc/internal/config"
	"github.com/think-free/sportsnews/providers/htafc/internal/upstream/model"
)

// Client implements the actual connection to the upstream service
// It implements the Upstream interface and can be used by the service

type Client struct {
	c *config.Config

	restClient *restylog.Client // I'm using a library that I've created to log the requests/response for another project that use 'resty' and respect the same interface
	logFilter  restylog.LogFilterFct
}

func New(ctx context.Context, c *config.Config) *Client {
	cli := &Client{
		c:          c,
		restClient: restylog.New(ctx),
		logFilter:  restylog.LogEverything,
	}

	cli.restClient.SetHostURL(cli.c.FeedURL)

	return cli
}

func (c *Client) GetNews(ctx context.Context) (*model.NewsArticles, error) {
	res := &model.NewsArticles{}
	urlPath := fmt.Sprintf("/getnewlistinformation?count=%d", c.c.ArticlesCount)
	err := c.getRequest(ctx, urlPath, res)
	return res, err
}

func (c *Client) GetNewsByID(ctx context.Context, id string) (*model.NewsArticle, error) {
	res := &model.NewsArticle{}
	urlPath := fmt.Sprintf("/getnewsarticleinformation?id=%s", id)
	err := c.getRequest(ctx, urlPath, res)
	return res, err
}

func (c *Client) getRequest(ctx context.Context, url string, res interface{}) error {
	r := c.restClient.R().
		SetContext(ctx).
		SetLogFilter(c.logFilter).
		SetResult(res)

	resp, err := r.Get(url)

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("error getting news, bad http status code: %d", resp.StatusCode())
	}

	return nil
}
