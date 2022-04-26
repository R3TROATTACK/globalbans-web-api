package server

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"insanitygaming.net/bans/src/gb/controllers/application"
	"insanitygaming.net/bans/src/gb/models/server"
	"insanitygaming.net/bans/src/gb/services/database"
)

func Find(app *gin.Context, id uint) (*server.Server, error) {
	var server server.Server
	database := app.MustGet("database").(*database.Database)
	row, err := database.QueryRow(context.Background(), "SELECT * FROM gb_server WHERE id = ?", id)
	if err != nil || row == nil {
		return nil, errors.New("Server not found")
	}
	var appid uint
	row.Scan(&server.Id, &appid, &server.Name, &server.Ip, &server.Port, &server.CreatedAt, &server.UpdatedAt, &server.Password)
	service, _ := application.Find(app, server.App.Id)
	server.App = *service
	return &server, nil
}

func FindByName(app *gin.Context, ip string, port uint) (*server.Server, error) {
	var server server.Server
	database := app.MustGet("database").(*database.Database)
	row, err := database.QueryRow(context.Background(), "SELECT * FROM gb_server WHERE ip = ? AND port = ?", ip, port)
	if err != nil || row == nil {
		return nil, errors.New("Server not found")
	}
	var appid uint
	row.Scan(&server.Id, &appid, &server.Name, &server.Ip, &server.Port, &server.CreatedAt, &server.UpdatedAt, &server.Password)
	service, _ := application.Find(app, server.App.Id)
	server.App = *service
	return &server, nil
}
