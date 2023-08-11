package database_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/think-free/sportsnews/sportsnews/internal/cliparams"
	"github.com/think-free/sportsnews/sportsnews/internal/config"
	"github.com/think-free/sportsnews/sportsnews/internal/database"
)

const (
	documenteID = "611834"
	teamID      = "Huddersfield Town"
)

func TestDatabase(t *testing.T) {
	ctx := context.Background()
	cp := &cliparams.ClientParameters{
		DatabaseURL: "mongodb://localhost:27017",
		Environment: "mock",
	}
	c := &config.Config{
		PageSize: 10,
	}

	if isManualTest() {
		cp.Environment = "production"
	}

	db := database.New(ctx, cp, c)

	testGetNewsByID(t, db)
	testGetAllNews(t, db)
	testGetPageNews(t, db)
}

func testGetNewsByID(t *testing.T, db database.Database) {
	article, err := db.GetNewsByID(context.Background(), teamID, documenteID)
	require.NoError(t, err)
	require.NotNil(t, article)
	require.Equal(t, documenteID, article.ID)
}

func testGetAllNews(t *testing.T, db database.Database) {
	articles, err := db.GetNews(context.Background(), teamID, -1)
	require.NoError(t, err)
	require.NotNil(t, articles)

	// This test is not completed, need to add more checks
}

func testGetPageNews(t *testing.T, db database.Database) {
	articles, err := db.GetNews(context.Background(), teamID, 4)
	require.NoError(t, err)
	require.NotNil(t, articles)

	// This test is not completed, need to add more checks
}

func isManualTest() bool {
	return os.Getenv("RUN_MANUAL_TESTS") == "true"
}
