package model

import (
	"context"
	"database/sql"
	"errors"

	"github.com/uptrace/bun"
)

type RoomRole string

const (
	RoomRoleSuspended     RoomRole = "Suspended"
	RoomRoleMember        RoomRole = "Member"
	RoomRoleModerator     RoomRole = "Moderator"
	RoomRoleAdministrator RoomRole = "Administrator"
)

type Room struct {
	bun.BaseModel `bun:"table:rooms"`

	ID          int    `bun:"id,pk,autoincrement" json:"id"`
	Name        string `bun:"name,notnull" json:"name"`
	Description string `bun:"description" json:"description"`
	Domain      string `bun:"domain,notnull" json:"domain"`
}

func CreateRoom(ctx context.Context, db *bun.DB, name, description, domain string) (*Room, error) {
	room := &Room{
		Name:        name,
		Description: description,
		Domain:      domain,
	}

	_, err := db.NewInsert().
		Model(room).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func GetAllRooms(ctx context.Context, db *bun.DB) (*[]Room, error) {
	var rooms []Room

	err := db.NewSelect().
		Model(&rooms).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &rooms, nil
}
