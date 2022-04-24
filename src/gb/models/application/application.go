package application

import (
	"context"

	"insanitygaming.net/bans/src/gb/services/database"
)

type Application struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (a *Application) Save(ctx context.Context) (bool, error) {
	_, err := database.Exec(ctx, "INSERT INTO gb_application (name, image) VALUES (?, ?)", a.Name, a.Image)
	return err == nil, err
}

func New(name string, image string) *Application {
	return &Application{
		Name:  name,
		Image: image,
	}
}
