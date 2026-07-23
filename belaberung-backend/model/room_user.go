package model

import (
	"context"
	"database/sql"
	"errors"

	"github.com/uptrace/bun"
)

type RoomUser struct {
	bun.BaseModel `bun:"table:room_users"`

	RoomID int `bun:",pk"`
	UserID int `bun:",pk"`

	Room *Room `bun:"rel:belongs-to,join:room_id=id"`
	User *User `bun:"rel:belongs-to,join:user_id=id"`

	Role RoomRole `bun:"role" json:"role"`
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

func GetRoomUserByIDs(ctx context.Context, db *bun.DB, roomID, userID int) (*RoomUser, error) {
	room_user := &RoomUser{
		RoomID: roomID,
		UserID: userID,
	}

	err := db.NewSelect().Model(room_user).WherePK().Relation("User").Relation("Room").Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return room_user, nil
}

func GetRoomUsersByRoomID(ctx context.Context, db *bun.DB, roomID int) ([]RoomUser, error) {
	var room_users []RoomUser

	err := db.NewSelect().Model(&room_users).Where("room_id = ?", roomID).Relation("User").Relation("Room").Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return room_users, nil
}

func GetRoomUsersByUserID(ctx context.Context, db *bun.DB, userID int) ([]RoomUser, error) {
	var room_users []RoomUser

	err := db.NewSelect().Model(&room_users).Where("user_id = ?", userID).Relation("User").Relation("Room").Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return room_users, nil
}

func GetRoomUsersByRoomIDAndRole(ctx context.Context, db *bun.DB, roomID int, role RoomRole) ([]RoomUser, error) {
	var room_users []RoomUser

	err := db.NewSelect().Model(&room_users).Where("room_id = ?", roomID).Where("role = ?", role).Relation("User").Relation("Room").Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return room_users, nil
}

func GetAllRoomsOfAnUser(ctx context.Context, db *bun.DB, userID int) ([]RoomUser, error) {
	var room_users []RoomUser

	err := db.NewSelect().Model(&room_users).Where("user_id = ?", userID).Relation("User").Relation("Room").Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return room_users, nil
}

func UpdateRoomUserRole(ctx context.Context, db *bun.DB, roomID, userID int, role RoomRole) (*RoomUser, error) {
	room_user, err := GetRoomUserByIDs(ctx, db, roomID, userID)
	if err != nil {
		return nil, err
	}
	if room_user == nil {
		return nil, nil
	}

	_, err = db.NewUpdate().Model(room_user).Set("role = ?", role).WherePK().Exec(ctx)

	if err != nil {
		return nil, err
	}

	return room_user, nil
}

func DeleteRoomUser(ctx context.Context, db *bun.DB, roomID, userID int) error {
	room_user := &RoomUser{
		RoomID: roomID,
		UserID: userID,
	}

	_, err := db.NewDelete().Model(room_user).WherePK().Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
