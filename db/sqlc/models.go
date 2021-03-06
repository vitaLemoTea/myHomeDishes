// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"database/sql"
)

type Account struct {
	ID           int32          `json:"id"`
	UserName     sql.NullString `json:"user_name"`
	PasswordHash sql.NullString `json:"password_hash"`
	CreatedAt    sql.NullTime   `json:"created_at"`
}

type Comment struct {
	ID          int32          `json:"id"`
	Content     sql.NullString `json:"content"`
	ReviewersID sql.NullInt32  `json:"reviewers_id"`
	CommentedID sql.NullInt32  `json:"commented_id"`
	CreatedAt   sql.NullTime   `json:"created_at"`
}

type People struct {
	ID              int32          `json:"id"`
	Name            sql.NullString `json:"name"`
	UploadAccountID sql.NullInt32  `json:"upload_account_id"`
	Desc            sql.NullString `json:"desc"`
	CreatedAt       sql.NullTime   `json:"created_at"`
}
