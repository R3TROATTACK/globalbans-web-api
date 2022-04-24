package server

import (
	"time"
)

type Group struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Servers   []uint    `json:"servers"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(name string, servers []uint) *Group {
	return &Group{
		Name:      name,
		Servers:   servers,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
