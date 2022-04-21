package server

import (
	"context"
	"errors"
	"time"

	"insanitygaming.net/bans/models/entity/application"
	"insanitygaming.net/bans/services/database"
)

type Server struct {
	Id        uint                    `json:"id" gorm:"primaryKey"`
	Name      string                  `json:"name"`
	Ip        string                  `json:"ip" gorm:"index:server_unique"`
	Port      uint                    `json:"port" gorm:"index:server_unique"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
	Password  string                  `json:"password"`
	App       application.Application `json:"application"`
	// Players   []Player `json:"players"`
}

func New(name string, ip string, port uint, password string, app application.Application) *Server {
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
	_, err := database.Exec(ctx, "INSERT INTO servers (app, name, ip, port, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)", s.App.Id, s.Name, s.Ip, s.Port, s.Password, s.CreatedAt, s.UpdatedAt)
	return err
}

func Find(ctx context.Context, id uint) (*Server, error) {
	var server Server
	row, err := database.QueryRow(ctx, "SELECT * FROM servers WHERE id = ?", id)
	if err != nil || row == nil {
		return nil, errors.New("Server not found")
	}
	var appid uint
	row.Scan(&server.Id, &appid, &server.Name, &server.Ip, &server.Port, &server.CreatedAt, &server.UpdatedAt, &server.Password)
	app, _ := application.Find(ctx, server.App.Id)
	server.App = *app
	return &server, nil
}

func FindByName(ctx context.Context, ip string, port uint) (*Server, error) {
	var server Server
	row, err := database.QueryRow(ctx, "SELECT * FROM servers WHERE ip = ? AND port = ?", ip, port)
	if err != nil || row == nil {
		return nil, errors.New("Server not found")
	}
	var appid uint
	row.Scan(&server.Id, &appid, &server.Name, &server.Ip, &server.Port, &server.CreatedAt, &server.UpdatedAt, &server.Password)
	app, _ := application.Find(ctx, server.App.Id)
	server.App = *app
	return &server, nil
}
