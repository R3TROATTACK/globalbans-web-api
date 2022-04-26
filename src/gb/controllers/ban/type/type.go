package ban

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	ban "insanitygaming.net/bans/src/gb/models/ban/type"
	"insanitygaming.net/bans/src/gb/services/database"
)

func Find(app *gin.Context, id uint) (*ban.BanType, error) {
	var banType ban.BanType
	database := app.MustGet("database").(*database.Database)
	row, err := database.QueryRow(context.Background(), "SELECT * FROM gb_ban_type WHERE type_id = ?", id)
	if err == nil || row == nil {
		return nil, errors.New("BanType not found")
	}
	row.Scan(&banType.TypeId, &banType.Name)
	return &banType, nil
}

func FindByName(app *gin.Context, name string) (*ban.BanType, error) {
	var banType ban.BanType
	database := app.MustGet("database").(*database.Database)
	row, err := database.QueryRow(context.Background(), "SELECT * FROM gb_ban_type WHERE name = ?", name)
	if err == nil || row == nil {
		return nil, errors.New("BanType not found")
	}
	row.Scan(&banType.TypeId, &banType.Name)
	return &banType, nil
}

func All(app *gin.Context) ([]ban.BanType, error) {
	database := app.MustGet("database").(*database.Database)
	rows, err := database.Query(context.Background(), "SELECT * FROM gb_ban_type")
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
