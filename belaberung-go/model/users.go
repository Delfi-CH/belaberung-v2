package model

import (
	"github.com/go-pg/pg/v10"
	"delfi.dev/belaberung-v2/crypt"
)

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

func CreateUser(db *pg.DB, username, domain, password string) (*User, error) {
	hash, err := crypt.EncryptPassword(password)
	if err != nil {
		return nil, err
	}
	user := &User{
		Username:       username,
		Domain:         domain,
		Password:       hash,
		ProfilePicture: "default",
		GlobalRole:     GlobalUserRoleMember,
	}

	_, err = db.Model(user).Insert()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetAllUsers(db *pg.DB) ([]User, error) {
	var users []User

	err := db.Model(&users).Select() 

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return users, nil
}

func GetUserById(db *pg.DB, id int) (*User, error) {
	user := &User{ID: id}
	err := db.Model(user).WherePK().Select()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func ValidateUserPassword(db *pg.DB, id int, password string) (bool, error) {
	user := &User{ID: id}
	err := db.Model(user).WherePK().Select()

	if err != nil {
		return false, err
	}

	passwordCorrect := crypt.CheckPassword(user.Password, password)

	return passwordCorrect, nil
}

func GetUserByUsername(db *pg.DB, username string) (*User, error) {
	user := &User{}
	err := db.Model(user).Where("username = ?", username).Select()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func GetUserByGlobalRole(db *pg.DB, role GlobalUserRole) (*User, error) {
	user := &User{}
	err := db.Model(user).Where("global_role = ?", role).Select()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}