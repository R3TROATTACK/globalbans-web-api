package admin

import (
	"context"
	"errors"
	"time"

	"insanitygaming.net/bans/services/database"
)

type Admin struct {
	Id           uint      `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	Flags        string    `json:"flags"`
	Servers      []uint    `json:"servers,omitempty" bson:"servers"`
	CreatedAt    time.Time `json:"created_at"`
	WebGroups    []uint    `json:"web_groups,omitempty"`
	ServerGroups []uint    `json:"server_groups,omitempty"`
	AdminGroup   []uint    `json:"admin_group,omitempty"`
}

func New(username string, password string, email string, flags string, servers []uint) *Admin {
	return &Admin{
		Username:  username,
		Password:  password,
		Email:     email,
		Flags:     flags,
		Servers:   servers,
		CreatedAt: time.Now(),
	}
}

func (admin *Admin) Save(ctx context.Context) error {
	_, err := database.Exec(ctx, "INSERT INTO admins (username, password, email, flags, servers, created_at) VALUES (?, ?, ?, ?, ?, ?)", admin.Username, admin.Password, admin.Email, admin.Flags, admin.Servers, admin.CreatedAt)
	return err
}

func Find(ctx context.Context, id uint) (*Admin, error) {
	var admin Admin
	row, err := database.QueryRow(ctx, "SELECT * FROM admins WHERE id = ?", id)
	if err != nil {
		return nil, errors.New("Admin not found")
	}
	row.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.Email, &admin.Flags, &admin.CreatedAt)
	return &admin, nil
}

func FindByName(ctx context.Context, username string) (*Admin, error) {
	var admin Admin
	row, err := database.QueryRow(ctx, "SELECT * FROM admins WHERE username = ?", username)
	if err != nil {
		return nil, errors.New("Admin not found")
	}
	row.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.Email, &admin.Flags, &admin.CreatedAt)
	return &admin, nil
}

func FindByServerId(ctx context.Context, id uint) ([]*Admin, error) {
	var admins []*Admin
	rows, err := database.Query(ctx, "SELECT * FROM admins WHERE servers = ?", id)
	if err != nil {
		return nil, errors.New("Admin not found")
	}
	for rows.Next() {
		var admin Admin
		rows.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.Email, &admin.Flags, &admin.CreatedAt)
		admins = append(admins, &admin)
	}
	return admins, nil
}

func FindByServerGroup(ctx context.Context, group uint) ([]*Admin, error) {
	var admins []*Admin
	rows, err := database.Query(ctx, "SELECT * FROM admins WHERE FIND_IN_SET(?, server_groups) = ?", group)
	if err != nil {
		return nil, errors.New("Admin not found")
	}
	for rows.Next() {
		var admin Admin
		rows.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.Email, &admin.Flags, &admin.CreatedAt)
		admins = append(admins, &admin)
	}
	return admins, nil
}
