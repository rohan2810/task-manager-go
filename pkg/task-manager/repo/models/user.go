package models

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64  `bun:"id,autoincrement" json:"id,omitempty"`
	Username      string `bun:"username,pk,notnull,unique" json:"username,omitempty"`
	Fullname      string `bun:"fullname,notnull" json:"fullname,omitempty"`
	Email         string `bun:"email,notnull" json:"email,omitempty"`
	Role          string `bun:"role,notnull" json:"role,omitempty"`
}
