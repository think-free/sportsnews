package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/think-free/sportsnews/lib/datamodel"
	"github.com/think-free/sportsnews/lib/icmongodb"
	"github.com/think-free/sportsnews/lib/logging"
	"github.com/think-free/sportsnews/providers/htafc/internal/cliparams"
)

type MongoDB struct {
	icmongodb.ICMongoDB
}

func New(ctx context.Context, cp *cliparams.ClientParameters) *MongoDB {
	return &MongoDB{
		ICMongoDB: *icmongodb.New(ctx, cp.DatabaseURL),
	}
}

func (m *MongoDB) SaveNews(ctx context.Context, news []datamodel.ICNews) error {
	// Creating a write operation for each article
	var operations []mongo.WriteModel
	for i := range news {
		article := news[i]
		if article.ID != "" {
			operations = append(operations,
				mongo.NewUpdateOneModel().
					SetFilter(bson.M{icmongodb.Key: article.ID}).
					SetUpdate(bson.M{"$set": article}).
					SetUpsert(true))
		} else {
			logging.L(ctx).Warnf("trying to insert article with empty id: %+v", article)
		}
	}

	// Executing a bulk write for all the operations (articles)
	result, err := m.GetCollection().BulkWrite(ctx, operations)
	if err != nil {
		return err
	}
	logging.L(ctx).Infof("inserted '%d' articles to database", len(result.UpsertedIDs))
	return nil
}
