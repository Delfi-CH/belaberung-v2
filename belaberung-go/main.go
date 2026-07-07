package main

import (
	"context"
	"delfi.dev/belaberung-v2/db"
	"delfi.dev/belaberung-v2/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"slices"
	"strconv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("No .env found...")
	}

	db, err := db.InitDB()
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
		role := c.Query("role")

		var users []model.User
		var err error

		if username == "" && role == "" {
			users, err = model.GetAllUsers(context.Background(), db)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
		} else if role == "" {
			var user *model.User
			user, err = model.GetUserByUsername(context.Background(), db, username)

			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}

			if user == nil {
				c.String(http.StatusNotFound, "user not found")
				return
			}

			users = []model.User{*user}
		} else if username == "" {
			users, err = model.GetUserByGlobalRole(context.Background(), db, model.GlobalUserRole(role))
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}

		} else {
			users, err = model.GetUserByGlobalRole(context.Background(), db, model.GlobalUserRole(role))
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			var user *model.User
			user, err = model.GetUserByUsername(context.Background(), db, username)

			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}

			if user == nil {
				c.String(http.StatusNotFound, "user not found")
				return
			}
			if slices.Contains(users, *user) {
				c.JSON(http.StatusOK, &user)
				return
			} else {
				c.String(http.StatusNotFound, "user not found")
				return
			}
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
		user, err := model.GetUserById(context.Background(), db, id)
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
