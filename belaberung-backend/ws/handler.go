package ws

import (
	"context"
	"log"
	"net/http"

	"delfi.dev/belaberung-v2/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func Handler(hub *Hub, db *bun.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionUsername := session.Get("username")
		if sessionUsername == nil {
			c.String(http.StatusUnauthorized, "not logged in")
			return
		}

		conn, err := Upgrader.Upgrade(
			c.Writer,
			c.Request,
			nil,
		)

		if err != nil {
			log.Println(err)
			return
		}


		client := newClient(
			hub,
			conn,
		)

		userID, err := model.GetUserByUsername(context.Background(), db, sessionUsername.(string))
		if err != nil {
			conn.Close()
			return
		}

		tmpRoomIDs, err := model.GetRoomUsersByUserID(context.Background(), db, userID.ID)

		roomIDs := []int{}

		for _, room := range tmpRoomIDs {
			roomIDs = append(roomIDs, room.RoomID)
		}

		if err != nil {
			conn.Close()
			return
		}

		for _, roomID := range roomIDs {

			hub.Subscribe(
				client,
				roomID,
			)

			client.rooms = append(
				client.rooms,
				roomID,
			)
		}


		go client.writePump()

		go client.readPump()
	}
}
