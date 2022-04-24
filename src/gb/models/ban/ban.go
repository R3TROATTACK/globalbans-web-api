package ban

import (
	"context"
	"time"

	"insanitygaming.net/bans/src/gb/models/admin"
	"insanitygaming.net/bans/src/gb/services/database"
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
	_, err := database.Exec(ctx, "INSERT INTO gb_bans (name, ip, auth_id, created_at, expires_at, reason, admin_id, server_id, comment, extra) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", ban.Name, ban.Ip, ban.AuthId, ban.CreatedAt, ban.ExpiresAt, ban.Reason, ban.Admin.Id, ban.ServerId, ban.Comment, ban.Extra)
	return err
}
