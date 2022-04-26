package application

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	model "insanitygaming.net/bans/src/gb/models/application"
	"insanitygaming.net/bans/src/gb/services/database"
)

func Find(app *gin.Context, id uint) (*model.Application, error) {
	var application model.Application
	database := app.MustGet("database").(*database.Database)
	row, err := database.QueryRow(context.Background(), "SELECT * FROM gb_application WHERE id = ?", id)
	if err != nil {
		return nil, errors.New("Application not found")
	}
	if row == nil {
		return nil, errors.New("Application not found")
	}
	row.Scan(&application.Id, &application.Name, &application.Image)

	return &application, nil
}

func FindByName(app *gin.Context, name string) (*model.Application, error) {
	var application model.Application
	database := app.MustGet("database").(*database.Database)
	row, err := database.QueryRow(context.Background(), "SELECT * FROM gb_application WHERE name = ?", name)
	if err != nil {
		return nil, errors.New("Application not found")
	}
	if row == nil {
		return nil, errors.New("Application not found")
	}
	row.Scan(&application.Id, &application.Name, &application.Image)

	return &application, nil
}
