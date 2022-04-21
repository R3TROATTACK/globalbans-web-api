package group

import (
	"context"
	"errors"
	"strconv"

	"insanitygaming.net/bans/services/database"
)

type Group struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Flags    uint   `json:"flags"`
	Immunity uint   `json:"immunity"`
}

func New(name string, permissions, flags uint, groupType uint, immunity uint) *Group {
	return &Group{
		Name:     name,
		Flags:    flags,
		Immunity: immunity,
	}
}

func Find(ctx context.Context, id uint) (*Group, error) {
	var Group Group

	rows, err := database.Query(ctx, "SELECT * FROM sb_groups WHERE group_id = ?", strconv.FormatUint(uint64(id), 10))
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("Group not found")
	}
	/*CREATE TABLE IF NOT EXISTS gb_groups (
		group_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(64) NOT NULL UNIQUE,
		flags INT NOT NULL DEFAULT 0,
		group_type INT UNSIGNED NOT NULL,
		immunity INT UNSIGNED NOT NULL DEFAULT 0
	);*/
	rows.Scan(&Group.Id, &Group.Name, &Group.Flags, &Group.Immunity)
	return &Group, nil
}

func FindByName(ctx context.Context, name string) (*Group, error) {
	var Group Group
	rows, err := database.Query(ctx, "SELECT * FROM sb_groups WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("Group not found")
	}
	rows.Scan(&Group.Id, &Group.Name, &Group.Flags, &Group.Immunity)
	return &Group, nil
}
