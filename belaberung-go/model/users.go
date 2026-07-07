package model

import (
	"context"
	"database/sql"
	"errors"

	"delfi.dev/belaberung-v2/crypt"
	"github.com/uptrace/bun"
)

type GlobalUserRole string

const (
	GlobalUserRoleMember                GlobalUserRole = "Member"
	GlobalUserRoleInstanceModerator     GlobalUserRole = "InstanceModerator"
	GlobalUserRoleInstanceAdministrator GlobalUserRole = "InstanceAdministrator"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID             int            `bun:"id,pk,autoincrement" json:"id"`
	Username       string         `bun:"username,unique,notnull" json:"username"`
	Biography      string         `bun:"biography" json:"biography"`
	ProfilePicture string         `bun:"profile_picture" json:"profilePicture"`
	Pronouns       string         `bun:"pronouns" json:"pronouns"`
	Domain         string         `bun:"domain,notnull" json:"domain"`
	Password       string         `bun:"password,notnull" json:"-"`
	Suspended      bool           `bun:"suspended" json:"suspended"`
	GlobalRole     GlobalUserRole `bun:"global_role" json:"globalRole"`
}

func CreateUser(ctx context.Context, db *bun.DB, username, domain, password string) (*User, error) {
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

	_, err = db.NewInsert().
		Model(user).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetAllUsers(ctx context.Context, db *bun.DB) ([]User, error) {
	var users []User

	err := db.NewSelect().
		Model(&users).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return users, nil
}

func GetUserById(ctx context.Context, db *bun.DB, id int) (*User, error) {
	user := &User{ID: id}

	err := db.NewSelect().
		Model(user).
		WherePK().
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func ValidateUserPassword(ctx context.Context, db *bun.DB, id int, password string) (bool, error) {
	user := &User{ID: id}

	err := db.NewSelect().
		Model(user).
		WherePK().
		Scan(ctx)

	if err != nil {
		return false, err
	}

	passwordCorrect := crypt.CheckPassword(user.Password, password)

	return passwordCorrect, nil
}

func GetUserByUsername(ctx context.Context, db *bun.DB, username string) (*User, error) {
	user := &User{}

	err := db.NewSelect().
		Model(user).
		Where("username = ?", username).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func GetUserByGlobalRole(ctx context.Context, db *bun.DB, role GlobalUserRole) ([]User, error) {
	var users []User

	err := db.NewSelect().
		Model(&users).
		Where("global_role = ?", role).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return users, nil
}

func UpdateUsername(ctx context.Context, db *bun.DB, id int, username string) (*User, error) {
	user := &User{ID: id}

	_, err := db.NewUpdate().
		Model(user).
		Set("username = ?", username).
		WherePK().
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}
