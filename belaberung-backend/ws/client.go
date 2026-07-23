package ws

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	hub *Hub

	conn *websocket.Conn

	send chan Message

	rooms []int
}


func newClient(hub *Hub, conn *websocket.Conn) *client {
	return &client{
		hub: hub,

		conn: conn,

		send: make(chan Message, 256),

		rooms: make([]int, 0),
	}
}


// readPump pumps messages from the websocket connection to the hub.
//
// There must only ever be one reader for a websocket connection.
func (c *client) readPump() {
	defer func() {

		// Remove this client from all rooms.
		for _, roomID := range c.rooms {
			c.hub.Unsubscribe(c, roomID)
		}

		c.conn.Close()
	}()


	c.conn.SetReadLimit(MaxMessageSize)

	c.conn.SetReadDeadline(
		time.Now().Add(PongWait),
	)


	c.conn.SetPongHandler(func(string) error {

		c.conn.SetReadDeadline(
			time.Now().Add(PongWait),
		)

		return nil
	})


	for {

		_, data, err := c.conn.ReadMessage()

		if err != nil {

			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				log.Println(err)
			}

			break
		}


		var message Message

		if err := json.Unmarshal(data, &message); err != nil {
			log.Println("invalid message:", err)
			continue
		}


		/*
		   IMPORTANT:

		   Usually you do NOT broadcast here.

		   You first save the message:

		       database.Insert(message)

		   Then:

		       hub.Broadcast(message)
		*/
		c.hub.Broadcast(message)
	}
}


// writePump pumps messages from the hub to the websocket connection.
//
// There must only ever be one writer for a websocket connection.
func (c *client) writePump() {

	ticker := time.NewTicker(PingPeriod)

	defer func() {

		ticker.Stop()

		c.conn.Close()
	}()


	for {

		select {


		case message, ok := <-c.send:


			c.conn.SetWriteDeadline(
				time.Now().Add(WriteWait),
			)


			if !ok {

				c.conn.WriteMessage(
					websocket.CloseMessage,
					nil,
				)

				return
			}


			data, err := json.Marshal(message)

			if err != nil {
				continue
			}


			err = c.conn.WriteMessage(
				websocket.TextMessage,
				data,
			)


			if err != nil {
				return
			}



		case <-ticker.C:


			c.conn.SetWriteDeadline(
				time.Now().Add(WriteWait),
			)


			if err := c.conn.WriteMessage(
				websocket.PingMessage,
				nil,
			); err != nil {
				return
			}
		}
	}
}


func (c *client) close() {

	for _, roomID := range c.rooms {

		c.hub.Unsubscribe(
			c,
			roomID,
		)
	}

	c.conn.Close()
}
