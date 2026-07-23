package route

import (
	"context"
	"net/http"
	"strconv"
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
				c.String(http.StatusNotFound, "room not found")
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
		password := c.Query("password")
		var room *model.Room
		if password == "" {
			room, err = model.GetRoomById(context.Background(), db, id)
		} else {
			room, err = model.GetPrivateRoomById(context.Background(), db, id, password)
		}

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if room == nil {
			c.String(http.StatusNotFound, "room not found")
			return
		}
		c.JSON(http.StatusOK, room)
	})

	//curl -b /tmp/cookies "http://localhost:8080/rooms/1"
	//curl -b /tmp/cookies "http://localhost:8080/rooms/1?password=1234"

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

		password := c.Query("password")
		var room *model.Room
		if password == "" {
			room, err = model.GetRoomById(context.Background(), db, id)
		} else {
			room, err = model.GetPrivateRoomById(context.Background(), db, id, password)
		}

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if room == nil {
			c.String(http.StatusNotFound, "room not found")
			return
		}

		var room_users []model.RoomUser

		role := c.Query("role")

		if role == "" {
			room_users, err = model.GetRoomUsersByRoomID(context.Background(), db, id)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			room_users, err = model.GetRoomUsersByRoomIDAndRole(context.Background(), db, id, model.RoomRole(role))
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
		}

		c.JSON(http.StatusOK, room_users)
	})

	//curl -b /tmp/cookies "http://localhost:8080/rooms/1/users"
	//curl -b /tmp/cookies "http://localhost:8080/rooms/1/users?password=1234"
	//curl -b /tmp/cookies "http://localhost:8080/rooms/1/users?role=Administrator"

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

		password := c.Query("password")
		var room *model.Room
		if password == "" {
			room, err = model.GetRoomById(context.Background(), db, id)
		} else {
			room, err = model.GetPrivateRoomById(context.Background(), db, id, password)
		}

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if room == nil {
			c.String(http.StatusNotFound, "room not found")
			return
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
	//curl -b /tmp/cookies "http://localhost:8080/rooms/1/join?password=1234"

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

		private := c.Query("private")
		var dbRoom *model.Room

		if private == "true" && room.Password != nil {
			dbRoom, err = model.CreatePrivateRoom(context.Background(), db, room.Name, room.Description, *room.Password)
		} else if private == "true" && room.Password == nil {
			c.String(http.StatusBadRequest, "private rooms need a password")
			return
		} else if private == "false" || private == "" {
			dbRoom, err = model.CreatePublicRoom(context.Background(), db, room.Name, room.Description)
		}

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
	//curl -X POST -H "Content-Type: application/json" -d '{"name":"private","description":"1234", "domain":"example.com", "password":"1234"}' -b /tmp/cookies "http://localhost:8080/rooms?private=true"

	router.DELETE("/:id", func(c *gin.Context) {
		session := sessions.Default(c)
		sessionUsername := session.Get("username")
		if sessionUsername == nil {
			c.String(http.StatusUnauthorized, "not logged in")
			return
		}

		idStr := c.Param("id")
		roomID, err := strconv.Atoi(idStr)
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

		password := c.Query("password")
		var room *model.Room
		if password == "" {
			room, err = model.GetRoomById(context.Background(), db, roomID)
		} else {
			room, err = model.GetPrivateRoomById(context.Background(), db, roomID, password)
		}

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if room == nil {
			c.String(http.StatusNotFound, "room not found")
			return
		}

		room_user, err := model.GetRoomUserByIDs(context.Background(), db, roomID, user.ID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		if room_user == nil {
			c.String(http.StatusNotFound, "user not found in room")
			return
		}

		if room_user.Role != "Administrator" {
			c.String(http.StatusForbidden, "permissions too low")
			return
		}

		err = model.DeleteRoom(context.Background(), db, roomID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusNoContent, "")
	})

	//curl -X DELETE -b /tmp/cookies "http://localhost:8080/rooms/1"
	//curl -X DELETE -b /tmp/cookies "http://localhost:8080/rooms/1?password=1234"
}
