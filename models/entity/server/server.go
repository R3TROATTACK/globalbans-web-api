package server

import (
	"context"
	"errors"
	"time"

	"insanitygaming.net/bans/services/database"
)

type Server struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Ip        string    `json:"ip" gorm:"index:server_unique"`
	Port      uint      `json:"port" gorm:"index:server_unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Password  string    `json:"password"`
	// Players   []Player `json:"players"`
}

func New(name string, ip string, port uint, password string) *Server {
	return &Server{
		Name:      name,
		Ip:        ip,
		Port:      port,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Password:  password,
	}
}

func Find(id uint) (*Server, error) {
	var server Server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := database.GetDatabase()
	res := db.Collection("servers").FindOne(ctx, &Server{Id: id})
	if res == nil {
		return nil, errors.New("Server not found")
	}
	res.Decode(&server)
	return &server, nil
}

func FindByName(ip string, port uint) (*Server, error) {
	var server Server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := database.GetDatabase()
	res := db.Collection("servers").FindOne(ctx, &Server{Ip: ip, Port: port})
	if res == nil {
		return nil, errors.New("Server not found")
	}
	res.Decode(&server)
	return &server, nil
}
