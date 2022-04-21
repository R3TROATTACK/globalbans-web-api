package Ban

import (
	"context"
	"errors"
	"time"

	"insanitygaming.net/bans/models/entity/admin"
	"insanitygaming.net/bans/services/database"
)

type Ban struct {
	BanId     uint        `json:"ban_id"`
	Name      string      `json:"name"`
	Ip        string      `json:"ip"`
	AuthId    uint        `json:"auth_id"`
	CreatedAt time.Time   `json:"created_at"`
	ExpiresAt time.Time   `json:"expires_at"`
	Reason    string      `json:"reason"`
	Admin     admin.Admin `json:"adminid"`
	ServerId  uint        `json:"serverid"`
	Comment   string      `json:"comment"`
	Extra     string      `json:"extra"`
}

func New(name string, ip string, authId uint, expiresAt time.Duration, reason string, adminId admin.Admin, serverId uint, comment string, extra string) *Ban {
	return &Ban{
		Name:      name,
		Ip:        ip,
		AuthId:    authId,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(expiresAt),
		Reason:    reason,
		Admin:     adminId,
		ServerId:  serverId,
		Comment:   comment,
		Extra:     extra,
	}
}

func (ban *Ban) Save(ctx context.Context) error {
	_, err := database.Exec(ctx, "INSERT INTO bans (name, ip, auth_id, created_at, expires_at, reason, admin_id, server_id, comment, extra) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", ban.Name, ban.Ip, ban.AuthId, ban.CreatedAt, ban.ExpiresAt, ban.Reason, ban.Admin.Id, ban.ServerId, ban.Comment, ban.Extra)
	return err
}

func Find(ctx context.Context, id uint) (*Ban, error) {
	var ban Ban
	row, err := database.QueryRow(ctx, "SELECT * FROM bans WHERE ban_id = ?", id)
	if err != nil || row == nil {
		return nil, errors.New("Ban not found")
	}
	var adminid uint
	row.Scan(&ban.BanId, &ban.Name, &ban.Ip, &ban.AuthId, &ban.CreatedAt, &ban.ExpiresAt, &ban.Reason, &adminid, &ban.ServerId, &ban.Comment, &ban.Extra)
	admin, _ := admin.Find(ctx, adminid)
	ban.Admin = *admin
	return &ban, nil
}

func FindByAuthId(ctx context.Context, authid uint) (*[]Ban, error) {
	var bans []Ban
	rows, err := database.Query(ctx, "SELECT * FROM bans WHERE player_id = ?", authid)
	if err != nil {
		return nil, errors.New("Ban not found")
	}
	defer rows.Close()
	for rows.Next() {
		var ban Ban
		var adminid uint
		rows.Scan(&ban.BanId, &ban.Name, &ban.Ip, &ban.AuthId, &ban.CreatedAt, &ban.ExpiresAt, &ban.Reason, &adminid, &ban.ServerId, &ban.Comment, &ban.Extra)
		admin, _ := admin.Find(ctx, adminid)
		ban.Admin = *admin
		bans = append(bans, ban)
	}
	return &bans, nil
}

func FindByServer(ctx context.Context, serverId uint) (*[]Ban, error) {
	var bans []Ban
	rows, err := database.Query(ctx, "SELECT * FROM bans WHERE server_id = ?", serverId)
	if err != nil {
		return nil, errors.New("Ban not found")
	}
	defer rows.Close()
	for rows.Next() {
		var ban Ban
		var adminid uint
		rows.Scan(&ban.BanId, &ban.Name, &ban.Ip, &ban.AuthId, &ban.CreatedAt, &ban.ExpiresAt, &ban.Reason, &adminid, &ban.ServerId, &ban.Comment, &ban.Extra)
		admin, _ := admin.Find(ctx, adminid)
		ban.Admin = *admin
		bans = append(bans, ban)
	}

	return &bans, nil
}
