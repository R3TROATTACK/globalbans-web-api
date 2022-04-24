package server

import (
	"context"
	"time"

	app "insanitygaming.net/bans/src/gb/models/application"
	"insanitygaming.net/bans/src/gb/services/database"
)

type Server struct {
	Id        uint            `json:"id" gorm:"primaryKey"`
	Name      string          `json:"name"`
	Ip        string          `json:"ip" gorm:"index:server_unique"`
	Port      uint            `json:"port" gorm:"index:server_unique"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Password  string          `json:"password"`
	App       app.Application `json:"application"`
	// Players   []Player `json:"players"`
}

func New(name string, ip string, port uint, password string, app app.Application) *Server {
	return &Server{
		Name:      name,
		Ip:        ip,
		Port:      port,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Password:  password,
		App:       app,
	}
}

func (s *Server) Save(ctx context.Context) error {
	_, err := database.Exec(ctx, "INSERT INTO gb_server (app, name, ip, port, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)", s.App.Id, s.Name, s.Ip, s.Port, s.Password, s.CreatedAt, s.UpdatedAt)
	return err
}
