package main

import (
	"fmt"
	"github.com/go-pg/pg/v10"
    "github.com/go-pg/pg/v10/orm"
	"delfi.dev/belaberung-v2/model"
	"net/http"
  	"github.com/gin-gonic/gin"
)

func main() {
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

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello, world")
	})

	r.GET("/users", func(c *gin.Context) {
		users, err := model.GetAllUsers(db)
		if err != nil {
			c.String(http.StatusInternalServerError, "error")
		}
		c.JSON(http.StatusOK, users)
	})
	
	r.Run()
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