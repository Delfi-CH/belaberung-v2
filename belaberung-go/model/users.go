package model

type GlobalUserRole string

const (
	GlobalUserRoleMember                GlobalUserRole = "Member"
	GlobalUserRoleInstanceModerator     GlobalUserRole = "InstanceModerator"
	GlobalUserRoleInstanceAdministrator GlobalUserRole = "InstanceAdministrator"
)

type User struct {
	tableName struct{} `pg:"users"`

	ID             int            `pg:"id,pk" json:"id"`
	Username       string         `pg:"username,unique,notnull" json:"username"`
	Biography      string         `pg:"biography" json:"biography"`
	ProfilePicture string         `pg:"profile_picture" json:"profilePicture"`
	Pronouns       string         `pg:"pronouns" json:"pronouns"`
	Domain         string         `pg:"domain,notnull" json:"domain"`
	Password       string         `pg:"password,notnull" json:"-"`
	Suspended      bool           `pg:"suspended,use_zero" json:"suspended"`
	GlobalRole     GlobalUserRole `pg:"global_role" json:"globalRole"`
}

func NewUser(username, domain, password string) *User {
	return &User{
		Username:       username,
		Domain:         domain,
		Password:       password,
		ProfilePicture: "default",
		GlobalRole:     GlobalUserRoleMember,
	}
}