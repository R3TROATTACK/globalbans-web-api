package web

import (
	"errors"
	"strconv"

	"insanitygaming.net/bans/src/gb"
	"insanitygaming.net/bans/src/gb/models/groups/web"
)

func Find(app *gb.GB, id uint) (*web.Group, error) {
	var webGroup web.Group

	rows, err := app.Database().Query(app.Context(), "SELECT group_id, name, flags, immunity FROM gb_group WHERE group_id = ? AND group_type = 0", strconv.FormatUint(uint64(id), 10))
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("WebGroup not found")
	}

	rows.Scan(&webGroup.Id, &webGroup.Name, &webGroup.Flags, &webGroup.Immunity)
	return &webGroup, nil
}

func FindByName(app *gb.GB, name string) (*web.Group, error) {
	var webGroup web.Group
	rows, err := app.Database().Query(app.Context(), "SELECT group_id, name, flags, immunity FROM gb_group WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("WebGroup not found")
	}
	rows.Scan(&webGroup.Id, &webGroup.Name, &webGroup.Flags, &webGroup.Immunity)
	return &webGroup, nil
}
