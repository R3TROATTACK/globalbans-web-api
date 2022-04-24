package application

import (
	"context"
	"errors"

	app "insanitygaming.net/bans/src/gb/models/application"
	"insanitygaming.net/bans/src/gb/services/database"
)

func Find(ctx context.Context, id uint) (*app.Application, error) {
	var application app.Application
	row, err := database.QueryRow(ctx, "SELECT * FROM gb_application WHERE id = ?", id)
	if err != nil {
		return nil, errors.New("Application not found")
	}
	if row == nil {
		return nil, errors.New("Application not found")
	}
	row.Scan(&application.Id, &application.Name, &application.Image)

	return &application, nil
}

func FindByName(ctx context.Context, name string) (*app.Application, error) {
	var application app.Application
	row, err := database.QueryRow(ctx, "SELECT * FROM gb_application WHERE name = ?", name)
	if err != nil {
		return nil, errors.New("Application not found")
	}
	if row == nil {
		return nil, errors.New("Application not found")
	}
	row.Scan(&application.Id, &application.Name, &application.Image)

	return &application, nil
}
