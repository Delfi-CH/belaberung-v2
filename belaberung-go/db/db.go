package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"delfi.dev/belaberung-v2/model"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func InitDB() (*bun.DB, error) {

	dbUser, exists := os.LookupEnv("BELABERUNG_DB_USERNAME")

	if !exists {
		fmt.Println("DB Username not set, using default...")
		dbUser = "belaberung"
	}

	dbPassword, exists := os.LookupEnv("BELABERUNG_DB_PASSWORD")

	if !exists {
		fmt.Println("DB Password not set, using default...")
		dbPassword = "belaberung"
	}

	dbName, exists := os.LookupEnv("BELABERUNG_DB_NAME")

	if !exists {
		fmt.Println("DB Name not set, using default...")
		dbName = "belaberung"
	}

	dbHost, exists := os.LookupEnv("BELABERUNG_DB_HOST")

	if !exists {
		fmt.Println("DB Host not set, using default...")
		dbHost = "localhost:5432"
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithAddr(dbHost),
		pgdriver.WithUser(dbUser),
		pgdriver.WithPassword(dbPassword),
		pgdriver.WithDatabase(dbName),
		pgdriver.WithInsecure(true),
	))

	db := bun.NewDB(sqldb, pgdialect.New())

	err := createSchema(context.Background(), db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createSchema(ctx context.Context, db *bun.DB) error {
	var err error

	_, err = db.NewCreateTable().
		Model((*model.User)(nil)).
		IfNotExists().
		Exec(ctx)

	if err != nil {
		return err
	}

	_, err = db.NewCreateTable().
		Model((*model.Room)(nil)).
		IfNotExists().
		Exec(ctx)

	if err != nil {
		return err
	}

	_, err = db.NewCreateTable().
		Model((*model.Message)(nil)).
		IfNotExists().
		WithForeignKeys().
		ForeignKey(`
		("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
	`).
		ForeignKey(`
		("room_id") REFERENCES "rooms" ("id") ON DELETE CASCADE
	`).
		Exec(ctx)

	if err != nil {
		return err
	}

	_, err = db.NewCreateTable().
		Model((*model.RoomUser)(nil)).
		IfNotExists().
		WithForeignKeys().
		ForeignKey(`
		("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
	`).
		ForeignKey(`
		("room_id") REFERENCES "rooms" ("id") ON DELETE CASCADE
	`).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
