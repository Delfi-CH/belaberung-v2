package model

import "github.com/go-pg/pg/v10"

type RoomUser struct {
	tableName struct{} `pg:"room_users"`

	RoomID int `pg:",pk,index:idx_room_user"`
	UserID int `pg:",pk,index:idx_room_user"`

	Role RoomRole
}

func CreateRoomUser(db *pg.DB, roomID, userID int) (*RoomUser, error) {
	room_user := &RoomUser{
		RoomID: roomID,
		UserID: userID,
		Role:   RoomRoleMember,
	}

	_, err := db.Model(room_user).Insert()
	if err != nil {
		
		return nil, err
	}

	return room_user, nil
}