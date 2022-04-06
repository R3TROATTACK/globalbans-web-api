package servergroup

import (
	"context"
	"errors"
	"time"

	"insanitygaming.net/bans/services/database"
)

type ServerGroup struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Servers   []string  `json:"servers"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(name string, servers []string) *ServerGroup {
	return &ServerGroup{
		Name:      name,
		Servers:   servers,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func Find(id uint) (*ServerGroup, error) {
	var serverGroup ServerGroup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := database.GetDatabase()
	res := db.Collection("servergroups").FindOne(ctx, &ServerGroup{Id: id})
	if res == nil {
		return nil, errors.New("ServerGroup not found")
	}
	res.Decode(&serverGroup)
	return &serverGroup, nil
}

func FindByName(name string) (*ServerGroup, error) {
	var serverGroup ServerGroup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := database.GetDatabase()
	res := db.Collection("servergroups").FindOne(ctx, &ServerGroup{Name: name})
	if res == nil {
		return nil, errors.New("ServerGroup not found")
	}
	res.Decode(&serverGroup)
	return &serverGroup, nil
}
