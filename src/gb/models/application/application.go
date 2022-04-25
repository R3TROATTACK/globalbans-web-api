package application

import (
	"insanitygaming.net/bans/src/gb"
)

type Application struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (a *Application) Save(app *gb.GB) (bool, error) {
	_, err := app.Database().Exec(app.Context(), "INSERT INTO gb_application (name, image) VALUES (?, ?)", a.Name, a.Image)
	return err == nil, err
}

func New(name string, image string) *Application {
	return &Application{
		Name:  name,
		Image: image,
	}
}
