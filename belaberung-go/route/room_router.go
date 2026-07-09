package route

import (
	"context"
	"strconv"
	"net/http"
	"strings"

	"delfi.dev/belaberung-v2/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func InitRoomRouter(router *gin.RouterGroup, db *bun.DB) {
	router.GET("", func(c *gin.Context) {
		session := sessions.Default(c)
		sessionUsername := session.Get("username")
		if sessionUsername == nil {
			c.String(http.StatusUnauthorized, "not logged in")
			return
		}
		roomname := c.Query("name")

		var err error
		var rooms []model.Room

		if roomname == "" {
			rooms, err = model.GetAllRooms(context.Background(), db)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			var room *model.Room
			room, err = model.GetRoomByName(context.Background(), db, roomname)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			if room == nil {
				c.String(http.StatusNotFound,"room not found")
				return
			}

			rooms = []model.Room{*room}
		}

		c.JSON(http.StatusOK, rooms)
	})

	//curl -b /tmp/cookies "http://localhost:8080/rooms"
	//curl -b /tmp/cookies "http://localhost:8080/rooms?name=demo"

	router.GET("/:id", func(c *gin.Context) {
		session := sessions.Default(c)
		sessionUsername := session.Get("username")
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
		room, err := model.GetRoomById(context.Background(), db, id)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, room)
	})

	//curl -b /tmp/cookies "http://localhost:8080/rooms/1"

	router.GET("/:id/users", func(c *gin.Context) {
		session := sessions.Default(c)
		sessionUsername := session.Get("username")
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
		room_users, err := model.GetRoomUsersByRoomID(context.Background(), db, id)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, room_users)
	})

	//curl -b /tmp/cookies "http://localhost:8080/rooms/1/users"

	router.GET("/:id/join", func(c *gin.Context) {
		session := sessions.Default(c)
		sessionUsername := session.Get("username")
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

		user, err := model.GetUserByUsername(context.Background(), db, sessionUsername.(string))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		if user == nil {
			c.String(http.StatusNotFound, "user not found")
		}

		room_user, err := model.CreateRoomUser(context.Background(), db, id, user.ID)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "room_users_key") || strings.Contains(err.Error(), "unique constraint") {
				c.String(http.StatusConflict, err.Error())
				return
			} else {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
		}
		c.JSON(http.StatusCreated, room_user)
	})

	//curl -b /tmp/cookies "http://localhost:8080/rooms/1/join"

	router.POST("", func(c *gin.Context) {
		session := sessions.Default(c)
		sessionUsername := session.Get("username")
		if sessionUsername == nil {
			c.String(http.StatusUnauthorized, "not logged in")
			return
		}

		var room model.CreateRoomRequest
		var err error 
		if err := c.ShouldBindJSON(&room); err != nil {
			c.String(http.StatusBadRequest, "bad request: "+err.Error())
			return
		}

		dbRoom, err := model.CreateRoom(context.Background(), db, room.Name, room.Description, room.Domain)

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "rooms_name_key") || strings.Contains(err.Error(), "unique constraint") {
				c.String(http.StatusConflict, err.Error())
				return
			} else {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
		}

		user, err := model.GetUserByUsername(context.Background(), db, sessionUsername.(string))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		if user == nil {
			c.String(http.StatusNotFound, "user not found")
			return
		}

		roomUser, err := model.CreateRoomUser(context.Background(), db, dbRoom.ID, user.ID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		roomUser, err = model.UpdateRoomUserRole(context.Background(), db, roomUser.RoomID, roomUser.UserID, model.RoomRoleAdministrator)

		c.JSON(http.StatusCreated, dbRoom)
	})

	//curl -X POST -H "Content-Type: application/json" -d '{"name":"test","description":"1234", "domain":"example.com"}' -b /tmp/cookies http://localhost:8080/rooms

	router.DELETE("/:id", func(c *gin.Context) {

	})
}