package ban

import (
	"errors"

	"insanitygaming.net/bans/src/gb"
	"insanitygaming.net/bans/src/gb/controllers/admin"
	"insanitygaming.net/bans/src/gb/models/ban"
)

func Find(app *gb.GB, id uint) (*ban.Ban, error) {
	var ban ban.Ban
	row, err := app.Database().QueryRow(app.Context(), "SELECT * FROM gb_bans WHERE ban_id = ?", id)
	if err != nil || row == nil {
		return nil, errors.New("Ban not found")
	}
	var adminid uint
	row.Scan(&ban.BanId, &ban.Name, &ban.Ip, &ban.AuthId, &ban.CreatedAt, &ban.ExpiresAt, &ban.Reason, &adminid, &ban.ServerId, &ban.Comment, &ban.Extra)
	admin, _ := admin.Find(app, adminid)
	ban.Admin = *admin
	return &ban, nil
}

func FindByAuthId(app *gb.GB, authid uint) (*[]ban.Ban, error) {
	var bans []ban.Ban
	rows, err := app.Database().Query(app.Context(), "SELECT * FROM gb_bans WHERE player_id = ?", authid)
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

func FindByServer(app *gb.GB, serverId uint) (*[]ban.Ban, error) {
	var bans []ban.Ban
	rows, err := app.Database().Query(app.Context(), "SELECT * FROM gb_bans WHERE server_id = ?", serverId)
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
