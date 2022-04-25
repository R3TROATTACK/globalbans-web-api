package admin

import (
	"time"

	"insanitygaming.net/bans/src/gb"
)

type Group struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Flags     string    `json:"flags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Immunity  uint      `json:"immunity"`
}

func (adminGroup *Group) Save(app *gb.GB) error {
	_, err := app.Database().Exec(app.Context(), "INSERT INTO gb_group (name, flags, immunity, group_type) VALUES (?, ?, ?, 1)", adminGroup.Name, adminGroup.Flags, adminGroup.Immunity)
	return err
}

func New(name string, flags string, immunity uint) *Group {
	return &Group{
		Name:      name,
		Flags:     flags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Immunity:  immunity,
	}
}
