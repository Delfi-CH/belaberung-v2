package ws

import "time"

type Message struct {
	ID         int       `json:"id"`
	Content    string    `json:"content"`
	Attachment any       `json:"attachment,omitempty"`
	Timestamp  time.Time `json:"timestamp"`
	UserID     int       `json:"user_id"`
	RoomID     int       `json:"room_id"`
}
