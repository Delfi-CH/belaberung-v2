package postgres

import (
	"github.com/go-pg/pg/v10"
    "github.com/go-pg/pg/v10/orm"
	"delfi.dev/belaberung-v2/model"
	"fmt"
	"os"
)

func InitDB() (*pg.DB, error) {

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

	db := pg.Connect(&pg.Options{
		Addr: dbHost,
		User: dbUser,
		Password: dbPassword,
		Database: dbName,
		ApplicationName: "belaberung",

	})

	err := createSchema(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*model.User)(nil),
		(*model.Room)(nil),
		(*model.RoomUser)(nil),
		(*model.Message)(nil),
		(*model.MessageAttachment)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}