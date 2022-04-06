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
	AdminId   admin.Admin `json:"admin"`
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
		AdminId:   adminId,
		ServerId:  serverId,
		Comment:   comment,
		Extra:     extra,
	}
}

func Find(id uint) (*Ban, error) {
	var ban Ban
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := database.GetDatabase()
	res := db.Collection("bans").FindOne(ctx, &Ban{BanId: id})
	if res == nil {
		return nil, errors.New("Ban not found")
	}
	res.Decode(&ban)
	return &ban, nil
}

func FindByAuthId(authid uint) (*[]Ban, error) {
	var ban []Ban
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := database.GetDatabase()
	cur, err := db.Collection("bans").Find(ctx, &Ban{AuthId: authid})
	if cur == nil || err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &ban); err != nil {
		return nil, err
	}
	return &ban, nil
}

func FindByServer(serverId uint) (*[]Ban, error) {
	var ban []Ban
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := database.GetDatabase()
	cur, err := db.Collection("bans").Find(ctx, &Ban{ServerId: serverId})
	if cur == nil || err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &ban); err != nil {
		return nil, err
	}
	return &ban, nil
}
