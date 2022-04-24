package ban

import (
	"context"

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

func (b *BanType) Save(ctx context.Context) error {
	_, err := database.Exec(ctx, "INSERT INTO gb_ban_type (name) VALUES (?)", b.Name)
	return err
}
