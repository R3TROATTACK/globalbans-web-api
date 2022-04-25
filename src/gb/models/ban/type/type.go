package ban

import "insanitygaming.net/bans/src/gb"

type BanType struct {
	TypeId uint   `json:"type_id"`
	Name   string `json:"name"`
}

func New(name string) *BanType {
	return &BanType{
		Name: name,
	}
}

func (b *BanType) Save(app *gb.GB) error {
	_, err := app.Database().Exec(app.Context(), "INSERT INTO gb_ban_type (name) VALUES (?)", b.Name)
	return err
}
