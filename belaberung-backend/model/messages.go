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

	Timestamp time.Time

	UserID int `bun:",notnull"`
	RoomID int `bun:",notnull"`

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

func GetRecentMessages(ctx context.Context, db *bun.DB, roomID int, limit int, cursor int) ([]Message, error) {
	var messages []Message

	query := db.NewSelect().
		Model(&messages).
		Where("room_id = ?", roomID).
		Order("id DESC").
		Limit(limit)

	if cursor > 0 {
		query = query.Where("id < ?", cursor)
	}

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

/*
	Load Messages
	messages, err := GetRecentMessages(ctx, db, roomID, 50, 0)

	Get oldest
	nextCursor := messages[len(messages)-1].ID
	olderMessages, err := GetRecentMessages(ctx, db, roomID, 50, nextCursor)
*/

func GetLastMessageID(ctx context.Context, db *bun.DB, roomID int) (int, error) {
	var message Message

	err := db.NewSelect().
		Model(&message).
		Column("id").
		Where("room_id = ?", roomID).
		Order("id DESC").
		Limit(1).
		Scan(ctx)

	if err != nil {
		return 0, err
	}

	return message.ID, nil
}
