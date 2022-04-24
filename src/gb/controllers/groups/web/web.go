package web

import (
	"context"
	"errors"
	"strconv"

	"insanitygaming.net/bans/src/gb/models/groups/web"
	"insanitygaming.net/bans/src/gb/services/database"
)

func Find(ctx context.Context, id uint) (*web.Group, error) {
	var webGroup web.Group

	rows, err := database.Query(ctx, "SELECT group_id, name, flags, immunity FROM gb_group WHERE group_id = ? AND group_type = 0", strconv.FormatUint(uint64(id), 10))
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("WebGroup not found")
	}

	rows.Scan(&webGroup.Id, &webGroup.Name, &webGroup.Flags, &webGroup.Immunity)
	return &webGroup, nil
}

func FindByName(ctx context.Context, name string) (*web.Group, error) {
	var webGroup web.Group
	rows, err := database.Query(ctx, "SELECT group_id, name, flags, immunity FROM gb_group WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("WebGroup not found")
	}
	rows.Scan(&webGroup.Id, &webGroup.Name, &webGroup.Flags, &webGroup.Immunity)
	return &webGroup, nil
}
