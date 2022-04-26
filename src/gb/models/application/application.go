package application

import (
	"context"

	"github.com/gin-gonic/gin"
	"insanitygaming.net/bans/src/gb/services/database"
)

type Application struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (a *Application) Save(app *gin.Context) (bool, error) {
	database := database.New()
	_, err := database.Exec(context.Background(), "INSERT INTO gb_application (name, image) VALUES (?, ?)", a.Name, a.Image)
	return err == nil, err
}

func New(name string, image string) *Application {
	return &Application{
		Name:  name,
		Image: image,
	}
}
