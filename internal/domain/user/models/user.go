package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID           pgtype.UUID        `db:"id"`
	Login        pgtype.Text        `db:"login"`
	Email        string             `db:"email"`
	PasswordHash pgtype.Text        `db:"password_hash"`
	FirstName    pgtype.Text        `db:"first_name"`
	LastName     pgtype.Text        `db:"last_name"`
	AvatarUrl    pgtype.Text        `db:"avatar_url"`
	IsVerified   pgtype.Bool        `db:"is_verified"`
	CreatedAt    pgtype.Timestamptz `db:"created_at"`
	UpdatedAt    pgtype.Timestamptz `db:"updated_at"`
}
