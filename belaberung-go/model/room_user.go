package model

import (
	"context"

	"github.com/uptrace/bun"
)

type RoomUser struct {
	bun.BaseModel `bun:"table:room_users"`

	RoomID int `bun:",pk,index:idx_room_user"`
	UserID int `bun:",pk,index:idx_room_user"`

	Role RoomRole
}

func CreateRoomUser(ctx context.Context, db *bun.DB, roomID, userID int) (*RoomUser, error) {
	room_user := &RoomUser{
		RoomID: roomID,
		UserID: userID,
		Role:   RoomRoleMember,
	}

	_, err := db.NewInsert().
		Model(room_user).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return room_user, nil
}
