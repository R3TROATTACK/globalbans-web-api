package admin

import (
	"context"
	"errors"
	"strconv"
	"strings"

	admcontrol "insanitygaming.net/bans/src/gb/controllers/groups/admin"
	servercontrol "insanitygaming.net/bans/src/gb/controllers/groups/server"
	webcontrol "insanitygaming.net/bans/src/gb/controllers/groups/web"
	"insanitygaming.net/bans/src/gb/models/admin"
	adm "insanitygaming.net/bans/src/gb/models/groups/admin"
	"insanitygaming.net/bans/src/gb/models/groups/server"
	"insanitygaming.net/bans/src/gb/models/groups/web"
	"insanitygaming.net/bans/src/gb/services/database"
)

func Find(ctx context.Context, id uint) (*admin.Admin, error) {
	var admin admin.Admin
	row, err := database.QueryRow(ctx, "SELECT admin_id, name, password, email, created_at, adm_groups, web_groups, svr_groups, flags, immunity FROM gb_admin WHERE id = ?", id)
	if err != nil {
		return nil, errors.New("Admin not found")
	}
	var admgroups, webgroups, svrgroups string

	row.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.Email, &admin.CreatedAt, &admgroups, &webgroups, &svrgroups, &admin.Flags, &admin.Immunity)

	admin.AdminGroup = parseAdminGropsFromList(ctx, parseIntsFromString(admgroups))
	admin.WebGroups = parseWebGroupsFromList(ctx, parseIntsFromString(webgroups))
	admin.ServerGroups = parseServerGroupsFromList(ctx, parseIntsFromString(svrgroups))

	return &admin, nil
}

func FindByName(ctx context.Context, username string) (*admin.Admin, error) {
	var admin admin.Admin
	row, err := database.QueryRow(ctx, "SELECT admin_id, name, password, email, auths, created_at, groups, flags, immunity FROM gb_admin WHERE username = ?", username)
	if err != nil {
		return nil, errors.New("Admin not found")
	}
	var admgroups, webgroups, svrgroups string

	row.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.Email, &admin.CreatedAt, &admgroups, &webgroups, &svrgroups, &admin.Flags, &admin.Immunity)

	admin.AdminGroup = parseAdminGropsFromList(ctx, parseIntsFromString(admgroups))
	admin.WebGroups = parseWebGroupsFromList(ctx, parseIntsFromString(webgroups))
	admin.ServerGroups = parseServerGroupsFromList(ctx, parseIntsFromString(svrgroups))
	return &admin, nil
}

func FindByServerId(ctx context.Context, id uint) ([]*admin.Admin, error) {
	var admins []*admin.Admin
	rows, err := database.Query(ctx, "SELECT * FROM admins WHERE servers = ?", id)
	if err != nil {
		return nil, errors.New("Admin not found")
	}
	for rows.Next() {
		var admin admin.Admin

		var admgroups, webgroups, svrgroups string

		rows.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.Email, &admin.CreatedAt, &admgroups, &webgroups, &svrgroups, &admin.Flags, &admin.Immunity)

		admin.AdminGroup = parseAdminGropsFromList(ctx, parseIntsFromString(admgroups))
		admin.WebGroups = parseWebGroupsFromList(ctx, parseIntsFromString(webgroups))
		admin.ServerGroups = parseServerGroupsFromList(ctx, parseIntsFromString(svrgroups))
		admins = append(admins, &admin)
	}
	return admins, nil
}

func FindByServerGroup(ctx context.Context, group uint) ([]*admin.Admin, error) {
	var admins []*admin.Admin
	rows, err := database.Query(ctx, "SELECT * FROM admins WHERE FIND_IN_SET(?, server_groups) = ?", group)
	if err != nil {
		return nil, errors.New("Admin not found")
	}
	for rows.Next() {
		var admin admin.Admin
		var admgroups, webgroups, svrgroups string

		rows.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.Email, &admin.CreatedAt, &admgroups, &webgroups, &svrgroups, &admin.Flags, &admin.Immunity)

		admin.AdminGroup = parseAdminGropsFromList(ctx, parseIntsFromString(admgroups))
		admin.WebGroups = parseWebGroupsFromList(ctx, parseIntsFromString(webgroups))
		admin.ServerGroups = parseServerGroupsFromList(ctx, parseIntsFromString(svrgroups))
		admins = append(admins, &admin)
	}
	return admins, nil
}

func parseIntsFromString(groups string) []uint {
	var ids []uint
	for _, group := range strings.Split(groups, ",") {
		id, err := strconv.ParseUint(group, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, uint(id))
	}
	return ids
}

func parseAdminGropsFromList(ctx context.Context, groups []uint) []adm.Group {
	var adminGroups []adm.Group
	for _, group := range groups {
		adminGroup, err := admcontrol.Find(ctx, uint(group))
		if err != nil {
			continue
		}
		adminGroups = append(adminGroups, *adminGroup)
	}
	return adminGroups
}

func parseWebGroupsFromList(ctx context.Context, groups []uint) []web.Group {
	var webGroups []web.Group
	for _, group := range groups {
		webGroup, err := webcontrol.Find(ctx, uint(group))
		if err != nil {
			continue
		}
		webGroups = append(webGroups, *webGroup)
	}
	return webGroups
}

func parseServerGroupsFromList(ctx context.Context, groups []uint) []server.Group {
	var serverGroups []server.Group
	for _, group := range groups {
		serverGroup, err := servercontrol.Find(ctx, uint(group))
		if err != nil {
			continue
		}
		serverGroups = append(serverGroups, *serverGroup)
	}
	return serverGroups
}
