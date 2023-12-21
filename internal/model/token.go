package model

type TokenData struct {
	UserID int `db:"user_id"`
	RoleID int `db:"role_id"`
}
