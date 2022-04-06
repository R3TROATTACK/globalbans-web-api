package admingroup

import (
	"context"
	"errors"
	"time"

	"insanitygaming.net/bans/services/database"
)

type AdminGroup struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Flags     string    `json:"flags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Immunity  uint      `json:"immunity"`
}

func New(name string, flags string, immunity uint) *AdminGroup {
	return &AdminGroup{
		Name:      name,
		Flags:     flags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Immunity:  immunity,
	}
}

func Find(id uint) (*AdminGroup, error) {
	var adminGroup AdminGroup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := database.GetDatabase()
	res := db.Collection("admingroups").FindOne(ctx, &AdminGroup{Id: id})
	if res == nil {
		return nil, errors.New("AdminGroup not found")
	}
	res.Decode(&adminGroup)
	return &adminGroup, nil
}

func FindByName(name string) (*AdminGroup, error) {
	var adminGroup AdminGroup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := database.GetDatabase()
	res := db.Collection("admingroups").FindOne(ctx, &AdminGroup{Name: name})
	if res == nil {
		return nil, errors.New("AdminGroup not found")
	}
	res.Decode(&adminGroup)
	return &adminGroup, nil
}
