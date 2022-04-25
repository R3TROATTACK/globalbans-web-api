package server

import (
	"errors"

	"insanitygaming.net/bans/src/gb"
	"insanitygaming.net/bans/src/gb/controllers/application"
	"insanitygaming.net/bans/src/gb/models/server"
)

func Find(app *gb.GB, id uint) (*server.Server, error) {
	var server server.Server
	row, err := app.Database().QueryRow(app.Context(), "SELECT * FROM gb_server WHERE id = ?", id)
	if err != nil || row == nil {
		return nil, errors.New("Server not found")
	}
	var appid uint
	row.Scan(&server.Id, &appid, &server.Name, &server.Ip, &server.Port, &server.CreatedAt, &server.UpdatedAt, &server.Password)
	service, _ := application.Find(app, server.App.Id)
	server.App = *service
	return &server, nil
}

func FindByName(app *gb.GB, ip string, port uint) (*server.Server, error) {
	var server server.Server
	row, err := app.Database().QueryRow(app.Context(), "SELECT * FROM gb_server WHERE ip = ? AND port = ?", ip, port)
	if err != nil || row == nil {
		return nil, errors.New("Server not found")
	}
	var appid uint
	row.Scan(&server.Id, &appid, &server.Name, &server.Ip, &server.Port, &server.CreatedAt, &server.UpdatedAt, &server.Password)
	service, _ := application.Find(app, server.App.Id)
	server.App = *service
	return &server, nil
}
