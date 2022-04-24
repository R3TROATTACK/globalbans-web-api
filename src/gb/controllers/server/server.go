package server

import (
	"context"
	"errors"

	app "insanitygaming.net/bans/src/gb/controllers/application"
	"insanitygaming.net/bans/src/gb/models/server"
	"insanitygaming.net/bans/src/gb/services/database"
)

func Find(ctx context.Context, id uint) (*server.Server, error) {
	var server server.Server
	row, err := database.QueryRow(ctx, "SELECT * FROM gb_server WHERE id = ?", id)
	if err != nil || row == nil {
		return nil, errors.New("Server not found")
	}
	var appid uint
	row.Scan(&server.Id, &appid, &server.Name, &server.Ip, &server.Port, &server.CreatedAt, &server.UpdatedAt, &server.Password)
	app, _ := app.Find(ctx, server.App.Id)
	server.App = *app
	return &server, nil
}

func FindByName(ctx context.Context, ip string, port uint) (*server.Server, error) {
	var server server.Server
	row, err := database.QueryRow(ctx, "SELECT * FROM gb_server WHERE ip = ? AND port = ?", ip, port)
	if err != nil || row == nil {
		return nil, errors.New("Server not found")
	}
	var appid uint
	row.Scan(&server.Id, &appid, &server.Name, &server.Ip, &server.Port, &server.CreatedAt, &server.UpdatedAt, &server.Password)
	app, _ := app.Find(ctx, server.App.Id)
	server.App = *app
	return &server, nil
}
