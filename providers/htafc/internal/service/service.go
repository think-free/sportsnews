package service

import (
	"context"
	"time"

	"github.com/think-free/sportsnews/lib/datamodel"
	"github.com/think-free/sportsnews/lib/logging"
	"github.com/think-free/sportsnews/providers/htafc/internal/config"
	"github.com/think-free/sportsnews/providers/htafc/internal/database"
	"github.com/think-free/sportsnews/providers/htafc/internal/upstream"
	"github.com/think-free/sportsnews/providers/htafc/internal/upstream/model"
)

type Service struct {
	c  *config.Config
	up upstream.Upstream
	db database.Database
}

func New(ctx context.Context, c *config.Config, up upstream.Upstream, db database.Database) *Service {
	s := &Service{
		c:  c,
		up: up,
		db: db,
	}

	return s
}

func (s *Service) Run(ctx context.Context) {
	for {
		articles := s.fetchAllArticles(ctx)
		inews := s.convertArticleToINews(ctx, articles)
		s.saveINews(ctx, inews)

		time.Sleep(s.c.PoolingInterval)
	}
}

// fetchAllArticles fetches all the news from the upstream and for each news it fetches the full articles and returns them
func (s *Service) fetchAllArticles(ctx context.Context) []*model.NewsArticle {
	news, err := s.up.GetNews(ctx)
	if err != nil {
		logging.L(ctx).Errorf("error getting news: %w", err)
		return []*model.NewsArticle{}
	}
	return s.fetchFullArticles(ctx, news)
}

// fetchFullArticles fetches the full articles for each news
func (s *Service) fetchFullArticles(ctx context.Context, news *model.NewsArticles) []*model.NewsArticle {
	articles := []*model.NewsArticle{}
	for _, article := range news.Items.Articles {
		fullArticle, err := s.up.GetNewsByID(ctx, article.ID)
		if err != nil {
			logging.L(ctx).Errorf("error getting full article '%s': %w", article.ID, err)
		}
		articles = append(articles, fullArticle)
	}
	return articles
}

// convertArticleToINews converts the news articles to the internal datamodel
func (s *Service) convertArticleToINews(ctx context.Context, articles []*model.NewsArticle) []datamodel.ICNews {
	news := make([]datamodel.ICNews, len(articles))
	for i := range articles {
		news[i] = *s.newsToDatamodel(ctx, articles[i])
	}
	return news
}

// newsToDatamodel converts a news article to the internal datamodel
func (s *Service) newsToDatamodel(ctx context.Context, article *model.NewsArticle) *datamodel.ICNews {
	return &datamodel.ICNews{
		ID:          article.Article.ID,
		TeamID:      s.c.TeamID,
		OptaMatchID: article.Article.OptaMatchID,
		Title:       article.Article.Title,
		Teaser:      article.Article.TeaserText,
		Content:     article.Article.BodyText,
		URL:         article.Article.ArticleURL,
		ImageURL:    article.Article.ThumbnailImageURL,
		GalleryUrls: article.Article.GalleryImageURLs,
		VideoURL:    article.Article.VideoURL,
		Published:   article.Article.PublishDate,
	}
}

// saveINews saves the news to the database
func (s *Service) saveINews(ctx context.Context, inews []datamodel.ICNews) {
	err := s.db.SaveNews(ctx, inews)
	if err != nil {
		logging.L(ctx).Errorf("error inserting news to database: %w", err)
	}
}
