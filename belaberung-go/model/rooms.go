package model

import "github.com/go-pg/pg/v10"

type RoomRole string

const (
	RoomRoleSuspended     RoomRole = "Suspended"
	RoomRoleMember        RoomRole = "Member"
	RoomRoleModerator     RoomRole = "Moderator"
	RoomRoleAdministrator RoomRole = "Administrator"
)

type Room struct {
	tableName struct{} `pg:"rooms"`

	ID          int    `pg:"id,pk" json:"id"`
	Name        string `pg:"name,notnull" json:"name"`
	Description string `pg:"description" json:"description"`
	Domain      string `pg:"domain,notnull" json:"domain"`
}

func CreateRoom(db *pg.DB, name, description, domain string) (*Room, error) {
	room := &Room{
		Name:        name,
		Description: description,
		Domain:      domain,
	}

	_, err := db.Model(room).Insert()
	if err != nil {
		return nil, err
	}

	return room, nil
}