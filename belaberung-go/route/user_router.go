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

	router.DELETE("/:id", func(c *gin.Context) {
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

		err = model.DeleteUser(context.Background(), db, id)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.String(http.StatusNoContent, "delete successfull")
	})

	//curl -X DELETE -b /tmp/cookies http://localhost:8080/users/1

	router.PATCH("/:id", func(c *gin.Context) {
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

		var req model.UpdateUserDetailRequest
		if err = c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "bad request: "+err.Error())
			return
		}

		detailType := req.RequestType

		switch detailType {
			case "username":
				if req.Username != nil {
					user, err = model.UpdateUserUsername(context.Background(), db, id, *req.Username)
					break
				} else {
					c.String(http.StatusBadRequest, "bad request: "+ err.Error())
					return
				}
			case "password":
				if req.OldPassword != nil || req.NewPassword != nil {
					user, err = model.UpdateUserPassword(context.Background(), db, id, *req.OldPassword, *req.NewPassword)
					break
				} else {
					c.String(http.StatusBadRequest, "bad request: "+ err.Error())
					return
				}
			case "biography":
				if req.Biography != nil {
					user, err = model.UpdateUserBiograpgy(context.Background(), db, id, *req.Biography)
					break
				} else {
					c.String(http.StatusBadRequest, "bad request: "+ err.Error())
					return
				}
			case "pronouns":
				if req.Pronouns != nil {
					user, err = model.UpdateUserPronouns(context.Background(), db, id, *req.Pronouns)
					break
				} else {
					c.String(http.StatusBadRequest, "bad request: "+ err.Error())
					return
				}
		}

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		
		c.JSON(http.StatusOK, &user)
	})

	//curl -X PATCH -H "Content-Type: application/json" -d '{"type":"pronouns","pronouns":"he/him"}' -b /tmp/cookies http://localhost:8080/users/1
}