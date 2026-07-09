package route

import (
	"context"
	"net/http"
	"slices"
	"strconv"

	"delfi.dev/belaberung-v2/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func InitUserRouter(router *gin.RouterGroup, db *bun.DB) {
	router.GET("", func(c *gin.Context) {

		session := sessions.Default(c)
		sessionUsername := session.Get("username")
		sessionIsAdministrator := session.Get("admin")

		if sessionUsername == nil {
			c.String(http.StatusUnauthorized, "not logged in")
			return
		}

		username := c.Query("name")
		role := c.Query("role")

		var users []model.User
		var err error

		if username == "" && role == "" && sessionIsAdministrator == true {
			users, err = model.GetAllUsers(context.Background(), db)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
		} else if role == "" && (username == sessionUsername || sessionIsAdministrator == true) {
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

	//curl -b /tmp/cookies http://localhost:8080/users
	//curl -b /tmp/cookies http://localhost:8080/users?name=demo
	//curl -b /tmp/cookies http://localhost:8080/users?role=Member
	//curl -b /tmp/cookies http://localhost:8080/users?name=demo&role=member

	router.GET("/:id", func(c *gin.Context) {
		session := sessions.Default(c)
		sessionUsername := session.Get("username")
		sessionIsAdministrator := session.Get("admin")

		if sessionUsername == nil {
			c.String(http.StatusUnauthorized, "not logged in")
			return
		}

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

		if user.Username != sessionUsername || sessionIsAdministrator == false {
			c.String(http.StatusForbidden, "invailid permissions")
			return
		}

		c.JSON(http.StatusOK, &user)
	})

	//curl -b /tmp/cookies http://localhost:8080/users/1
}
