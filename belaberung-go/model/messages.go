package model

import (
	"context"
	"encoding/json"
	"time"

	"github.com/uptrace/bun"
)

type MessageAttachmentType string

const (
	MessageAttachmentNone       MessageAttachmentType = "None"
	MessageAttachmentMention    MessageAttachmentType = "Mention"
	MessageAttachmentReply      MessageAttachmentType = "Reply"
	MessageAttachmentBinaryFile MessageAttachmentType = "BinaryFile"
	MessageAttachmentTextFile   MessageAttachmentType = "TextFile"
	MessageAttachmentDocument   MessageAttachmentType = "Document"
	MessageAttachmentHyperlink  MessageAttachmentType = "Hyperlink"
	MessageAttachmentImage      MessageAttachmentType = "Image"
	MessageAttachmentAudio      MessageAttachmentType = "Audio"
	MessageAttachmentVideo      MessageAttachmentType = "Video"
)

type MessageAttachment struct {
	Type MessageAttachmentType `json:"type"`
	Data json.RawMessage       `json:"data"`
}

type Message struct {
	bun.BaseModel `bun:"table:messages"`

	ID int `bun:"id,pk,autoincrement"`

	Content string

	Attachment MessageAttachment `bun:",type:jsonb"`

	Timestamp time.Time `bun:",index:idx_room_time"`

	UserID int `bun:",notnull,index:idx_user_messages"`
	RoomID int `bun:",notnull,index:idx_room_time"`

	User *User `bun:"rel:belongs-to,join:user_id=id"`
	Room *Room `bun:"rel:belongs-to,join:room_id=id"`
}

func CreateMessage(ctx context.Context, db *bun.DB, content string, attachment MessageAttachment, userID, roomID int) (*Message, error) {
	message := &Message{
		Content:    content,
		Attachment: attachment,
		Timestamp:  time.Now(),
		UserID:     userID,
		RoomID:     roomID,
	}

	_, err := db.NewInsert().
		Model(message).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return message, nil
}
