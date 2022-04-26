package ban

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"insanitygaming.net/bans/src/gb/controllers/admin"
	"insanitygaming.net/bans/src/gb/models/ban"
	"insanitygaming.net/bans/src/gb/services/database"
)

func Find(app *gin.Context, id uint) (*ban.Ban, error) {
	var ban ban.Ban
	database := app.MustGet("database").(*database.Database)
	row, err := database.QueryRow(context.Background(), "SELECT * FROM gb_bans WHERE ban_id = ?", id)
	if err != nil || row == nil {
		return nil, errors.New("Ban not found")
	}
	var adminid uint
	row.Scan(&ban.BanId, &ban.Name, &ban.Ip, &ban.AuthId, &ban.CreatedAt, &ban.ExpiresAt, &ban.Reason, &adminid, &ban.ServerId, &ban.Comment, &ban.Extra)
	admin, _ := admin.Find(app, adminid)
	ban.Admin = *admin
	return &ban, nil
}

func FindByAuthId(app *gin.Context, authid uint) (*[]ban.Ban, error) {
	var bans []ban.Ban
	database := app.MustGet("database").(*database.Database)
	rows, err := database.Query(context.Background(), "SELECT * FROM gb_bans WHERE player_id = ?", authid)
	if err != nil {
		return nil, errors.New("Ban not found")
	}
	defer rows.Close()
	for rows.Next() {
		var ban ban.Ban
		var adminid uint
		rows.Scan(&ban.BanId, &ban.Name, &ban.Ip, &ban.AuthId, &ban.CreatedAt, &ban.ExpiresAt, &ban.Reason, &adminid, &ban.ServerId, &ban.Comment, &ban.Extra)
		admin, _ := admin.Find(app, adminid)
		ban.Admin = *admin
		bans = append(bans, ban)
	}
	return &bans, nil
}

func FindByServer(app *gin.Context, serverId uint) (*[]ban.Ban, error) {
	var bans []ban.Ban
	database := app.MustGet("database").(*database.Database)
	rows, err := database.Query(context.Background(), "SELECT * FROM gb_bans WHERE server_id = ?", serverId)
	if err != nil {
		return nil, errors.New("Ban not found")
	}
	defer rows.Close()
	for rows.Next() {
		var ban ban.Ban
		var adminid uint
		rows.Scan(&ban.BanId, &ban.Name, &ban.Ip, &ban.AuthId, &ban.CreatedAt, &ban.ExpiresAt, &ban.Reason, &adminid, &ban.ServerId, &ban.Comment, &ban.Extra)
		admin, _ := admin.Find(app, adminid)
		ban.Admin = *admin
		bans = append(bans, ban)
	}

	return &bans, nil
}
