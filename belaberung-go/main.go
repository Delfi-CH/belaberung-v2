package main

import (
	"fmt"
	"github.com/go-pg/pg/v10"
    "github.com/go-pg/pg/v10/orm"
	"delfi.dev/belaberung-v2/model"
)

func main() {
	fmt.Println("hello, world")

	db := pg.Connect(&pg.Options{
		User: "belaberung",
		Password: "belaberung",
		Database: "belaberung",
	})
	defer db.Close()

	err := createSchema(db)
	if err != nil {
		fmt.Println("schema error")
		panic(err)
	}

	demoUser := model.NewUser("demo", "example.com", "1234")

	_, err = db.Model(demoUser).Insert()
	if err != nil {
		fmt.Println("insert error")
		panic(err)
	}

	var users []model.User
	err = db.Model(&users).Select()
	if err != nil {
		fmt.Println("select error")
		panic(err)
	}
	fmt.Println(users)
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
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}