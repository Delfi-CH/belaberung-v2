package main

import (
	"delfi.dev/belaberung-v2/model"
	"delfi.dev/belaberung-v2/postgres"
	"github.com/joho/godotenv"
	"net/http"
	"fmt"
	"strconv"
  	"github.com/gin-gonic/gin"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("No .env found...")
	}

	db, err := postgres.InitDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello, world")
	})

	r.GET("/users", func(c *gin.Context) {
		username := c.Query("name")
		
		var users []model.User
		var err error

		if username == "" {
			users, err = model.GetAllUsers(db)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			var user *model.User
			user, err = model.GetUserByUsername(db, username)
			
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}

			if user == nil {
				c.String(http.StatusNotFound, "user not found")
				return
			}

			users = []model.User{*user}
		}

		c.JSON(http.StatusOK, &users)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		user, err := model.GetUserById(db, id)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		
		if user == nil {
			c.String(http.StatusNotFound, "user not found")
			return
		}

		c.JSON(http.StatusOK, &user)
	})
	
	r.Run()
}