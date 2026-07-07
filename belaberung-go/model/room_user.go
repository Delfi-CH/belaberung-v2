package model

type RoomUser struct {
	tableName struct{} `pg:"room_users"`

	RoomID int `pg:",pk,index:idx_room_user"`
	UserID int `pg:",pk,index:idx_room_user"`

	Role RoomRole
}

func NewRoomUser(roomID, userID int) *RoomUser {
	return &RoomUser{
		RoomID: roomID,
		UserID: userID,
		Role:   RoomRoleMember,
	}
}