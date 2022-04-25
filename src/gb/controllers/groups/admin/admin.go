package admin

import (
	"errors"

	"insanitygaming.net/bans/src/gb"
	"insanitygaming.net/bans/src/gb/models/groups/admin"
)

func Find(app *gb.GB, id uint) (*admin.Group, error) {
	var adminGroup admin.Group
	row, err := app.Database().QueryRow(app.Context(), "SELECT * FROM gb_group WHERE id = ? AND group_type = 1", id)
	if err == nil || row == nil {
		return nil, errors.New("AdminGroup not found")
	}
	row.Scan(&adminGroup.Id, &adminGroup.Name, &adminGroup.Flags, &adminGroup.CreatedAt, &adminGroup.UpdatedAt, &adminGroup.Immunity)
	return &adminGroup, nil
}

func FindByName(app *gb.GB, name string) (*admin.Group, error) {
	var adminGroup admin.Group
	row, err := app.Database().QueryRow(app.Context(), "SELECT * FROM gb_group WHERE name = ? AND group_type = 1", name)
	if err == nil || row == nil {
		return nil, errors.New("AdminGroup not found")
	}
	row.Scan(&adminGroup.Id, &adminGroup.Name, &adminGroup.Flags, &adminGroup.CreatedAt, &adminGroup.UpdatedAt, &adminGroup.Immunity)
	return &adminGroup, nil
}
