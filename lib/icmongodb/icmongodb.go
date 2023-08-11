package icmongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/think-free/sportsnews/lib/logging"
)

const (
	DatabaseName   = "incrowd"
	CollectionName = "articles"
	Key            = "id"
)

type ICMongoDB struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func New(ctx context.Context, databaseURL string) *ICMongoDB {
	m := &ICMongoDB{}

	// Connecting to mongodb
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURL))
	if err != nil {
		logging.L(ctx).Fatalf("error connecting to mongodb: %w", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		logging.L(ctx).Fatalf("error pinging mongodb: %w", err)
	}

	// Creating collection and index
	coll := m.createCollection(ctx, client)
	err = m.createIndex(ctx, coll)
	if err != nil {
		logging.L(ctx).Fatalf("error creating index for id field: %w", err)
	}

	logging.L(ctx).Infof("connected to mongodb")

	m.client = client
	m.coll = coll
	return m
}

func (m *ICMongoDB) GetClient() *mongo.Client {
	return m.client
}

func (m *ICMongoDB) GetCollection() *mongo.Collection {
	return m.coll
}

func (m *ICMongoDB) createCollection(ctx context.Context, client *mongo.Client) *mongo.Collection {
	col := client.Database(DatabaseName).Collection(CollectionName)
	return col
}

func (m *ICMongoDB) createIndex(ctx context.Context, col *mongo.Collection) error {
	_, err := col.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    map[string]int{Key: 1},
		Options: options.Index().SetUnique(true),
	})
	return err
}
