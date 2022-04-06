package admin

import (
	"context"
	"errors"
	"time"

	admingroup "insanitygaming.net/bans/models/groups/admin"
	servergroup "insanitygaming.net/bans/models/groups/server"
	webgroup "insanitygaming.net/bans/models/groups/web"
	"insanitygaming.net/bans/services/database"
)

type Admin struct {
	Id           uint                      `json:"id"`
	Username     string                    `json:"username"`
	Password     string                    `json:"password"`
	Email        string                    `json:"email"`
	Flags        string                    `json:"flags"`
	Servers      []string                  `json:"servers,omitempty"`
	CreatedAt    time.Time                 `json:"created_at"`
	WebGroups    []webgroup.WebGroup       `json:"web_groups,omitempty"`
	ServerGroups []servergroup.ServerGroup `json:"server_groups,omitempty"`
	AdminGroup   []admingroup.AdminGroup   `json:"admin_group,omitempty"`
}

func New(username string, password string, email string, flags string, servers []string) *Admin {
	return &Admin{
		Username:  username,
		Password:  password,
		Email:     email,
		Flags:     flags,
		Servers:   servers,
		CreatedAt: time.Now(),
	}
}

func Find(id uint) (*Admin, error) {
	var admin Admin
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := database.GetDatabase()
	res := db.Collection("admins").FindOne(ctx, &Admin{Id: id})
	if res == nil {
		return nil, errors.New("Admin not found")
	}
	res.Decode(&admin)
	return &admin, nil
}

func FindByName(username string) (*Admin, error) {
	var admin Admin
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := database.GetDatabase()
	res := db.Collection("admins").FindOne(ctx, &Admin{Username: username})
	if res == nil {
		return nil, errors.New("Admin not found")
	}
	res.Decode(&admin)
	return &admin, nil
}
