package application

import (
	"context"
	"errors"

	"insanitygaming.net/bans/services/database"
)

type Application struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (a *Application) Save(ctx context.Context) (bool, error) {
	_, err := database.Exec(ctx, "INSERT INTO applications (name, image) VALUES (?, ?)", a.Name, a.Image)
	return err == nil, err
}

func New(name string, image string) *Application {
	return &Application{
		Name:  name,
		Image: image,
	}
}

func Find(ctx context.Context, id uint) (*Application, error) {
	var application Application
	row, err := database.QueryRow(ctx, "SELECT * FROM applications WHERE id = ?", id)
	if err != nil {
		return nil, errors.New("Application not found")
	}
	if row == nil {
		return nil, errors.New("Application not found")
	}
	row.Scan(&application.Id, &application.Name, &application.Image)

	return &application, nil
}

func FindByName(ctx context.Context, name string) (*Application, error) {
	var application Application
	row, err := database.QueryRow(ctx, "SELECT * FROM applications WHERE name = ?", name)
	if err != nil {
		return nil, errors.New("Application not found")
	}
	if row == nil {
		return nil, errors.New("Application not found")
	}
	row.Scan(&application.Id, &application.Name, &application.Image)

	return &application, nil
}
