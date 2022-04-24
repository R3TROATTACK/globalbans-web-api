package ban

import (
	"context"
	"errors"

	"insanitygaming.net/bans/src/gb/controllers/admin"
	"insanitygaming.net/bans/src/gb/models/ban"
	"insanitygaming.net/bans/src/gb/services/database"
)

func Find(ctx context.Context, id uint) (*ban.Ban, error) {
	var ban ban.Ban
	row, err := database.QueryRow(ctx, "SELECT * FROM gb_bans WHERE ban_id = ?", id)
	if err != nil || row == nil {
		return nil, errors.New("Ban not found")
	}
	var adminid uint
	row.Scan(&ban.BanId, &ban.Name, &ban.Ip, &ban.AuthId, &ban.CreatedAt, &ban.ExpiresAt, &ban.Reason, &adminid, &ban.ServerId, &ban.Comment, &ban.Extra)
	admin, _ := admin.Find(ctx, adminid)
	ban.Admin = *admin
	return &ban, nil
}

func FindByAuthId(ctx context.Context, authid uint) (*[]ban.Ban, error) {
	var bans []ban.Ban
	rows, err := database.Query(ctx, "SELECT * FROM gb_bans WHERE player_id = ?", authid)
	if err != nil {
		return nil, errors.New("Ban not found")
	}
	defer rows.Close()
	for rows.Next() {
		var ban ban.Ban
		var adminid uint
		rows.Scan(&ban.BanId, &ban.Name, &ban.Ip, &ban.AuthId, &ban.CreatedAt, &ban.ExpiresAt, &ban.Reason, &adminid, &ban.ServerId, &ban.Comment, &ban.Extra)
		admin, _ := admin.Find(ctx, adminid)
		ban.Admin = *admin
		bans = append(bans, ban)
	}
	return &bans, nil
}

func FindByServer(ctx context.Context, serverId uint) (*[]ban.Ban, error) {
	var bans []ban.Ban
	rows, err := database.Query(ctx, "SELECT * FROM gb_bans WHERE server_id = ?", serverId)
	if err != nil {
		return nil, errors.New("Ban not found")
	}
	defer rows.Close()
	for rows.Next() {
		var ban ban.Ban
		var adminid uint
		rows.Scan(&ban.BanId, &ban.Name, &ban.Ip, &ban.AuthId, &ban.CreatedAt, &ban.ExpiresAt, &ban.Reason, &adminid, &ban.ServerId, &ban.Comment, &ban.Extra)
		admin, _ := admin.Find(ctx, adminid)
		ban.Admin = *admin
		bans = append(bans, ban)
	}

	return &bans, nil
}
