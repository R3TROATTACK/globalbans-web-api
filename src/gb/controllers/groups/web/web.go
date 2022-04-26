package web

import (
	"context"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"insanitygaming.net/bans/src/gb/models/groups/web"
	"insanitygaming.net/bans/src/gb/services/database"
)

func Find(app *gin.Context, id uint) (*web.Group, error) {
	var webGroup web.Group
	database := app.MustGet("database").(*database.Database)
	rows, err := database.Query(context.Background(), "SELECT group_id, name, flags, immunity FROM gb_group WHERE group_id = ? AND group_type = 0", strconv.FormatUint(uint64(id), 10))
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("WebGroup not found")
	}

	rows.Scan(&webGroup.Id, &webGroup.Name, &webGroup.Flags, &webGroup.Immunity)
	return &webGroup, nil
}

func FindByName(app *gin.Context, name string) (*web.Group, error) {
	var webGroup web.Group
	database := app.MustGet("database").(*database.Database)
	rows, err := database.Query(context.Background(), "SELECT group_id, name, flags, immunity FROM gb_group WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("WebGroup not found")
	}
	rows.Scan(&webGroup.Id, &webGroup.Name, &webGroup.Flags, &webGroup.Immunity)
	return &webGroup, nil
}
