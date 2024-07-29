package inventory

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/sangharsh/devx-pipeline/inventory/db"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestInventory(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx,
		"postgres:16.3-alpine3.20",
		postgres.WithInitScripts(filepath.Join("..", "db-init.sql")),
		postgres.WithDatabase("inventory_db"),
		postgres.WithUsername("inventory_user"),
		postgres.WithPassword("abc123xyz"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(50*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)

	err = db.InitDB(ctx, connStr)
	assert.NoError(t, err)

	c, err := CreateInventory(Inventory{
		ProductName: "Cars",
		Quantity:    3,
	})
	assert.NoError(t, err)
	assert.NotNil(t, c)

	list, err := ListInventory()
	assert.NoError(t, err)
	assert.NotNil(t, list)
}
