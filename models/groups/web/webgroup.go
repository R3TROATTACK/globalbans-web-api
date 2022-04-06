package webgroup

import (
	"context"
	"errors"
	"time"

	"insanitygaming.net/bans/services/database"
)

type WebGroup struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Permissions []string  `json:"flags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Immunity    uint      `json:"immunity"`
}

func New(name string, permissions []string, immunity uint) *WebGroup {
	return &WebGroup{
		Name:        name,
		Permissions: permissions,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Immunity:    immunity,
	}
}

func Find(id uint) (*WebGroup, error) {
	var webGroup WebGroup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := database.GetDatabase()
	res := db.Collection("webgroups").FindOne(ctx, &WebGroup{Id: id})
	if res == nil {
		return nil, errors.New("WebGroup not found")
	}
	res.Decode(&webGroup)
	return &webGroup, nil
}

func FindByName(name string) (*WebGroup, error) {
	var webGroup WebGroup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := database.GetDatabase()
	res := db.Collection("webgroups").FindOne(ctx, &WebGroup{Name: name})
	if res == nil {
		return nil, errors.New("WebGroup not found")
	}
	res.Decode(&webGroup)
	return &webGroup, nil
}
