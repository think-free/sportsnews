package mongodb

import (
	"context"
	"fmt"

	"github.com/think-free/sportsnews/lib/datamodel"
	"github.com/think-free/sportsnews/lib/icmongodb"
	"github.com/think-free/sportsnews/lib/logging"
	"github.com/think-free/sportsnews/sportsnews/internal/cliparams"
	"github.com/think-free/sportsnews/sportsnews/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	icmongodb.ICMongoDB
	cp *cliparams.ClientParameters
	c  *config.Config
}

func New(ctx context.Context, cp *cliparams.ClientParameters, c *config.Config) *MongoDB {
	return &MongoDB{
		ICMongoDB: *icmongodb.New(ctx, cp.DatabaseURL),
		cp:        cp,
		c:         c,
	}
}

func (m *MongoDB) GetNews(ctx context.Context, team string, page int) ([]*datamodel.ICNews, error) {
	// Getting options and filters
	filter := m.getTeamFilter(ctx, team)
	ctx, opts := m.getPageOptions(ctx, page)

	// Getting articles
	var articles []*datamodel.ICNews
	cursor, err := m.GetCollection().Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("error getting page '%d' of articles: %s", page, err)
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &articles)
	if err != nil {
		return nil, fmt.Errorf("error decoding page '%d' of articles: %s", page, err)
	}

	// Returning articles
	logging.L(ctx).Infof("'%d' articles retreived from database", len(articles))
	return articles, nil
}

func (m *MongoDB) GetNewsByID(ctx context.Context, team string, id string) (*datamodel.ICNews, error) {
	// Getting filters
	filter := m.getTeamFilter(ctx, team)
	filter[icmongodb.Key] = id

	// Getting article
	var article datamodel.ICNews
	result := m.GetCollection().FindOne(ctx, filter)
	if result.Err() != nil {
		return nil, fmt.Errorf("error getting article with id '%s': %s", id, result.Err())
	}
	err := result.Decode(&article)
	if err != nil {
		return nil, fmt.Errorf("error decoding article with id '%s': %s", id, err)
	}

	// Returning article
	logging.L(ctx).Debugf("got article with id '%s': %+v", id, article)
	return &article, nil
}

// getPageOptions returns the FindOptions options for pagination, calculating the skip and limit values
func (m *MongoDB) getPageOptions(ctx context.Context, page int) (context.Context, *options.FindOptions) {
	opts := options.Find().SetSort(bson.D{{Key: "published", Value: -1}})
	ctx = logging.SetTag(ctx, "page.paging", false)
	if page >= 0 {
		ctx = logging.SetTag(ctx, "page.paging", true)
		ctx = logging.SetTag(ctx, "page.size", m.c.PageSize)
		ctx = logging.SetTag(ctx, "page.number", page)

		opts = opts.SetLimit(m.c.PageSize)

		if page > 0 {
			skip := int64(page * int(m.c.PageSize))
			ctx = logging.SetTag(ctx, "page.skipped_entries", skip)
			opts.SetSkip(skip)
		}
	}
	return ctx, opts
}

// getTeamFilter returns the filter for the team if it is specified, if the team is empty it returns an empty filter and all the teams will be returned
func (m *MongoDB) getTeamFilter(ctx context.Context, team string) bson.M {
	if team == "" {
		return bson.M{}
	}
	return bson.M{"teamId": team}
}
