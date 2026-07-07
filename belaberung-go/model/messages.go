package model

import "github.com/go-pg/pg/v10"

import (
	"encoding/json"
	"time"
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
	tableName struct{} `pg:"messages"`

	ID int `pg:"id,pk"`

	Content string

	Attachment MessageAttachment `pg:"type:jsonb"`

	Timestamp time.Time `pg:",index:idx_room_time"`

	UserID int `pg:",notnull,index:idx_user_messages"`
	RoomID int `pg:",notnull,index:idx_room_time"`

	User *User `pg:"rel:has-one"`
	Room *Room `pg:"rel:has-one"`
}

func CreateMessage(db *pg.DB, content string, attachment MessageAttachment, userID, roomID int) (*Message, error) {
	message := &Message{
		Content:    content,
		Attachment: attachment,
		Timestamp:  time.Now(),
		UserID:     userID,
		RoomID:     roomID,
	}

	_, err := db.Model(message).Insert()
	if err != nil {
		return nil, err
	}

	return message, nil
}