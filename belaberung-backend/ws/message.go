package ws

import (
	"context"
	"github.com/uptrace/bun"
	"time"

	"delfi.dev/belaberung-v2/model"
)

type Message struct {
	ID         int                     `json:"id"`
	Content    string                  `json:"content"`
	Attachment model.MessageAttachment `json:"attachment,omitempty"`
	Timestamp  time.Time               `json:"timestamp"`
	UserID     int                     `json:"user_id"`
	Username   string                  `json:"username"`
	Role       model.RoomRole          `json:"role"`
	RoomID     int                     `json:"room_id"`
}

func FromDatabaseMessage(message *model.Message, db *bun.DB) (*Message, error) {
	roomUser, err := model.GetRoomUserByIDs(context.Background(), db, message.RoomID, message.UserID)
	if err != nil {
		return nil, err
	}
	return &Message{
		ID:         message.ID,
		Content:    message.Content,
		Attachment: message.Attachment,
		Timestamp:  message.Timestamp,
		UserID:     message.UserID,
		Role:       roomUser.Role,
		Username:   message.User.Username,
		RoomID:     message.RoomID,
	}, nil
}
