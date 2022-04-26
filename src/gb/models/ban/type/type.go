package ban

import (
	"context"

	"github.com/gin-gonic/gin"
	"insanitygaming.net/bans/src/gb/services/database"
)

type BanType struct {
	TypeId uint   `json:"type_id"`
	Name   string `json:"name"`
}

func New(name string) *BanType {
	return &BanType{
		Name: name,
	}
}

func (b *BanType) Save(app *gin.Context) error {
	database := database.New()
	_, err := database.Exec(context.Background(), "INSERT INTO gb_ban_type (name) VALUES (?)", b.Name)
	return err
}
