package model

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

func NewRoom(name, description, domain string) *Room {
	return &Room{
		Name:        name,
		Description: description,
		Domain:      domain,
	}
}