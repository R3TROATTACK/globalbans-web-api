package admin

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"insanitygaming.net/bans/src/gb/models/groups/admin"
	"insanitygaming.net/bans/src/gb/services/database"
)

func Find(app *gin.Context, id uint) (*admin.Group, error) {
	var adminGroup admin.Group
	database := app.MustGet("database").(*database.Database)
	row, err := database.QueryRow(context.Background(), "SELECT * FROM gb_group WHERE id = ? AND group_type = 1", id)
	if err == nil || row == nil {
		return nil, errors.New("AdminGroup not found")
	}
	row.Scan(&adminGroup.Id, &adminGroup.Name, &adminGroup.Flags, &adminGroup.CreatedAt, &adminGroup.UpdatedAt, &adminGroup.Immunity)
	return &adminGroup, nil
}

func FindByName(app *gin.Context, name string) (*admin.Group, error) {
	var adminGroup admin.Group
	database := app.MustGet("database").(*database.Database)
	row, err := database.QueryRow(context.Background(), "SELECT * FROM gb_group WHERE name = ? AND group_type = 1", name)
	if err == nil || row == nil {
		return nil, errors.New("AdminGroup not found")
	}
	row.Scan(&adminGroup.Id, &adminGroup.Name, &adminGroup.Flags, &adminGroup.CreatedAt, &adminGroup.UpdatedAt, &adminGroup.Immunity)
	return &adminGroup, nil
}
