package web

import (
	"context"
	"time"

	"insanitygaming.net/bans/src/gb/services/database"
)

type Group struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Flags     string    `json:"flags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Immunity  uint      `json:"immunity"`
}

func New(name string, permissions, flags string, immunity uint) *Group {
	return &Group{
		Name:      name,
		Flags:     flags,
		CreatedAt: time.Now(),
		Immunity:  immunity,
	}
}

func (webGroup *Group) Save(ctx context.Context) error {
	_, err := database.Exec(ctx, "INSERT INTO gb_group (name, flags, immunity) VALUES (?, ?, ?)", webGroup.Name, webGroup.Flags, webGroup.Immunity)
	return err
}
