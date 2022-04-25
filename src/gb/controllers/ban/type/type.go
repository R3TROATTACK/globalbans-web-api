package ban

import (
	"errors"

	"insanitygaming.net/bans/src/gb"
	ban "insanitygaming.net/bans/src/gb/models/ban/type"
)

func Find(app *gb.GB, id uint) (*ban.BanType, error) {
	var banType ban.BanType
	row, err := app.Database().QueryRow(app.Context(), "SELECT * FROM gb_ban_type WHERE type_id = ?", id)
	if err == nil || row == nil {
		return nil, errors.New("BanType not found")
	}
	row.Scan(&banType.TypeId, &banType.Name)
	return &banType, nil
}

func FindByName(app *gb.GB, name string) (*ban.BanType, error) {
	var banType ban.BanType
	row, err := app.Database().QueryRow(app.Context(), "SELECT * FROM gb_ban_type WHERE name = ?", name)
	if err == nil || row == nil {
		return nil, errors.New("BanType not found")
	}
	row.Scan(&banType.TypeId, &banType.Name)
	return &banType, nil
}

func All(app *gb.GB) ([]ban.BanType, error) {
	rows, err := app.Database().Query(app.Context(), "SELECT * FROM gb_ban_type")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var banTypes []ban.BanType
	for rows.Next() {
		var banType ban.BanType
		err := rows.Scan(&banType.TypeId, &banType.Name)
		if err != nil {
			return nil, err
		}
		banTypes = append(banTypes, banType)
	}
	return banTypes, nil
}
