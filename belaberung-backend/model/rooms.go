package model

import (
	"context"
	"database/sql"
	"errors"

	"delfi.dev/belaberung-v2/crypt"
	"github.com/uptrace/bun"
	"os"
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
	Name        string `bun:"name,unique,notnull" json:"name"`
	Description string `bun:"description" json:"description"`
	Domain      string `bun:"domain,notnull" json:"domain"`
	IsPrivate   bool   `bun:"private,notnull" json:"private"`
	Password    string `bun:"password" json:"-"`
}

func CreatePublicRoom(ctx context.Context, db *bun.DB, name, description string) (*Room, error) {
	domain, exists := os.LookupEnv("BELABERUNG_DOMAIN")

	if !exists {
		domain = "example.com"
	}

	room := &Room{
		Name:        name,
		Description: description,
		Domain:      domain,
		IsPrivate:   false,
	}

	_, err := db.NewInsert().
		Model(room).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func CreatePrivateRoom(ctx context.Context, db *bun.DB, name, description, password string) (*Room, error) {
	domain, exists := os.LookupEnv("BELABERUNG_DOMAIN")

	if !exists {
		domain = "example.com"
	}
	hash, err := crypt.EncryptPassword(password)
	if err != nil {
		return nil, err
	}

	room := &Room{
		Name:        name,
		Description: description,
		Domain:      domain,
		IsPrivate:   true,
		Password:    hash,
	}

	_, err = db.NewInsert().
		Model(room).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func GetAllRooms(ctx context.Context, db *bun.DB) ([]Room, error) {
	var rooms []Room

	err := db.NewSelect().
		Model(&rooms).
		Where("private = ?", false).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return rooms, nil
}

func GetRoomById(ctx context.Context, db *bun.DB, id int) (*Room, error) {
	room := &Room{ID: id}
	err := db.NewSelect().Model(room).WherePK().Where("private = ?", false).Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return room, nil
}

func GetPrivateRoomById(ctx context.Context, db *bun.DB, id int, password string) (*Room, error) {
	room := &Room{ID: id}
	err := db.NewSelect().Model(room).WherePK().Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if crypt.CheckPassword(room.Password, password) {
		return room, nil
	} else {
		return nil, nil
	}
}

func GetRoomByName(ctx context.Context, db *bun.DB, name string) (*Room, error) {
	room := &Room{}
	err := db.NewSelect().Model(room).Where("name = ?", name).Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return room, nil
}

func UpdateRoomName(ctx context.Context, db *bun.DB, id int, name string) (*Room, error) {
	room := &Room{ID: id}

	_, err := db.NewUpdate().
		Model(room).
		Set("name = ?", name).
		WherePK().
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return room, nil
}

func UpdateRoomDescription(ctx context.Context, db *bun.DB, id int, description string) (*Room, error) {
	room := &Room{ID: id}

	_, err := db.NewUpdate().
		Model(room).
		Set("description = ?", description).
		WherePK().
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return room, nil
}

func DeleteRoom(ctx context.Context, db *bun.DB, id int) error {
	room := &Room{ID: id}

	_, err := db.NewDelete().Model(room).WherePK().Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
