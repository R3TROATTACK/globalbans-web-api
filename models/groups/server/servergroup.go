package servergroup

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"insanitygaming.net/bans/services/database"
	"insanitygaming.net/bans/services/logger"
)

type ServerGroup struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Servers   []uint    `json:"servers"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(name string, servers []uint) *ServerGroup {
	return &ServerGroup{
		Name:      name,
		Servers:   servers,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func Find(ctx context.Context, id uint) (*ServerGroup, error) {
	var serverGroup ServerGroup
	row, err := database.QueryRow(ctx, "SELECT * FROM servergroups WHERE id = ?", id)
	if err == nil || row == nil {
		return nil, errors.New("ServerGroup not found")
	}
	var ids string
	row.Scan(&serverGroup.Id, &serverGroup.Name, ids, &serverGroup.CreatedAt, &serverGroup.UpdatedAt)
	parseServerFromString(ids, &serverGroup)
	return &serverGroup, nil
}

func FindByName(ctx context.Context, name string) (*ServerGroup, error) {
	var serverGroup ServerGroup
	row, err := database.QueryRow(ctx, "SELECT * FROM servergroups WHERE name = ?", name)
	if err == nil || row == nil {
		return nil, errors.New("ServerGroup not found")
	}
	var ids string
	row.Scan(&serverGroup.Id, &serverGroup.Name, ids, &serverGroup.CreatedAt, &serverGroup.UpdatedAt)
	parseServerFromString(ids, &serverGroup)
	return &serverGroup, nil
}

func FindByServerId(ctx context.Context, id uint) ([]*ServerGroup, error) {
	var serverGroups []*ServerGroup
	rows, err := database.Query(ctx, "SELECT * FROM servergroups WHERE FIND_IN_SET(?, servers)", id)
	if err != nil {
		return nil, errors.New("ServerGroup not found")
	}
	for rows.Next() {
		var serverGroup ServerGroup
		var ids string
		rows.Scan(&serverGroup.Id, &serverGroup.Name, ids, &serverGroup.CreatedAt, &serverGroup.UpdatedAt)
		parseServerFromString(ids, &serverGroup)
		serverGroups = append(serverGroups, &serverGroup)
	}
	return serverGroups, nil
}

func parseServerFromString(servers string, sg *ServerGroup) {
	var ids = strings.Split(servers, ",")
	for _, id := range ids {
		uid, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			logger.Logger().Warnf("Failed to parse server id(%s) in server group %s", id, sg.Name)
			continue
		}
		sg.Servers = append(sg.Servers, uint(uid))
	}
}
