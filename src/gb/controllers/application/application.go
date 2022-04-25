package application

import (
	"errors"

	"insanitygaming.net/bans/src/gb"
	model "insanitygaming.net/bans/src/gb/models/application"
)

func Find(app *gb.GB, id uint) (*model.Application, error) {
	var application model.Application
	row, err := app.Database().QueryRow(app.Context(), "SELECT * FROM gb_application WHERE id = ?", id)
	if err != nil {
		return nil, errors.New("Application not found")
	}
	if row == nil {
		return nil, errors.New("Application not found")
	}
	row.Scan(&application.Id, &application.Name, &application.Image)

	return &application, nil
}

func FindByName(app *gb.GB, name string) (*model.Application, error) {
	var application model.Application
	row, err := app.Database().QueryRow(app.Context(), "SELECT * FROM gb_application WHERE name = ?", name)
	if err != nil {
		return nil, errors.New("Application not found")
	}
	if row == nil {
		return nil, errors.New("Application not found")
	}
	row.Scan(&application.Id, &application.Name, &application.Image)

	return &application, nil
}
