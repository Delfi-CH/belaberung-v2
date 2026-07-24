package ws

import (
	"time"

	"delfi.dev/belaberung-v2/model"
)

type Message struct {
	ID         int       `json:"id"`
	Content    string    `json:"content"`
	Attachment model.MessageAttachment       `json:"attachment,omitempty"`
	Timestamp  time.Time `json:"timestamp"`
	UserID     int       `json:"user_id"`
	Username   string    `json:"username"`
	RoomID     int       `json:"room_id"`
}

func FromDatabaseMessage(message *model.Message) *Message {
	return &Message{
		ID: message.ID,
		Content: message.Content,
		Attachment: message.Attachment,
		Timestamp: message.Timestamp,
		UserID: message.UserID,
		Username: message.User.Username,
		RoomID: message.RoomID,
	}
}