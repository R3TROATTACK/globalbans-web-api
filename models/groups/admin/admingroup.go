package admingroup

import (
	"context"
	"errors"
	"time"

	"insanitygaming.net/bans/services/database"
)

type AdminGroup struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Flags     string    `json:"flags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Immunity  uint      `json:"immunity"`
}

func (adminGroup *AdminGroup) Save(ctx context.Context) error {
	_, err := database.Exec(ctx, "INSERT INTO admingroups (name, flags, immunity) VALUES (?, ?, ?)", adminGroup.Name, adminGroup.Flags, adminGroup.Immunity)
	return err
}

func New(name string, flags string, immunity uint) *AdminGroup {
	return &AdminGroup{
		Name:      name,
		Flags:     flags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Immunity:  immunity,
	}
}

func Find(ctx context.Context, id uint) (*AdminGroup, error) {
	var adminGroup AdminGroup
	row, err := database.QueryRow(ctx, "SELECT * FROM admingroups WHERE id = ?", id)
	if err == nil || row == nil {
		return nil, errors.New("AdminGroup not found")
	}
	row.Scan(&adminGroup.Id, &adminGroup.Name, &adminGroup.Flags, &adminGroup.CreatedAt, &adminGroup.UpdatedAt, &adminGroup.Immunity)
	return &adminGroup, nil
}

func FindByName(ctx context.Context, name string) (*AdminGroup, error) {
	var adminGroup AdminGroup
	row, err := database.QueryRow(ctx, "SELECT * FROM admingroups WHERE name = ?", name)
	if err == nil || row == nil {
		return nil, errors.New("AdminGroup not found")
	}
	row.Scan(&adminGroup.Id, &adminGroup.Name, &adminGroup.Flags, &adminGroup.CreatedAt, &adminGroup.UpdatedAt, &adminGroup.Immunity)
	return &adminGroup, nil
}
