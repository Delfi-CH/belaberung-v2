package route

import (
	"context"
	"net/http"

	"delfi.dev/belaberung-v2/crypt"
	"delfi.dev/belaberung-v2/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)


func InitAuthRouter(router *gin.RouterGroup, db *bun.DB) {
	router.POST("/login", func(c *gin.Context) {
		var user model.LoginRequest
		if err := c.ShouldBindJSON(&user); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}

		dbUser, err := model.GetUserByUsername(context.Background(), db, user.Username)
		if err != nil {
			c.String(http.StatusInternalServerError, "database error")
			return
		}

		if dbUser == nil {
			c.String(http.StatusNotFound, "user not found")
			return
		}

		if !crypt.CheckPassword(dbUser.Password, user.Password) {
			c.String(http.StatusUnauthorized, "false password")
			return
		}

		session := sessions.Default(c)
		session.Set("username", user.Username)
		err = session.Save()
		if err != nil {
			c.String(http.StatusInternalServerError, "cookie error")
			return
		}
		c.String(http.StatusOK, "login")
	})

	router.GET("/status", func(c *gin.Context) {
		session := sessions.Default(c)
		sessionUsername := session.Get("username")

		if sessionUsername != nil {
			c.String(http.StatusOK, "logged in")
			return
		} else {
			c.String(http.StatusForbidden, "not logged in")
		}
	})

	router.GET("/logout", func (c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()	
		err := session.Save()

		if err != nil {
			c.String(http.StatusInternalServerError, "logout failed")
			return
		}
		c.String(http.StatusOK, "logout sucessfull")
	})
}