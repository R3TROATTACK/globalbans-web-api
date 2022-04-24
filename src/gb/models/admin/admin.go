package admin

import (
	"context"
	"encoding/json"
	"sort"
	"strconv"
	"strings"
	"time"

	adm "insanitygaming.net/bans/src/gb/models/groups/admin"
	"insanitygaming.net/bans/src/gb/models/groups/server"
	"insanitygaming.net/bans/src/gb/models/groups/web"
	"insanitygaming.net/bans/src/gb/services/database"
)

type Admin struct {
	Id           uint              `json:"id"`
	Username     string            `json:"username"`
	Password     string            `json:"password"`
	Email        string            `json:"email"`
	Auths        map[string]string `json:"auths"`
	Flags        string            `json:"flags"`
	Immunity     uint              `json:"immunity"`
	Servers      []uint            `json:"servers,omitempty" bson:"servers"`
	CreatedAt    time.Time         `json:"created_at"`
	WebGroups    []web.Group       `json:"web_groups,omitempty"`
	ServerGroups []server.Group    `json:"server_groups,omitempty"`
	AdminGroup   []adm.Group       `json:"admin_group,omitempty"`
}

func New(username string, password string, email string, auths map[string]string, flags string, servers []uint, immunity uint) *Admin {
	return &Admin{
		Username:  username,
		Password:  password,
		Email:     email,
		Flags:     flags,
		Servers:   servers,
		Immunity:  immunity,
		CreatedAt: time.Now(),
		Auths:     auths,
	}
}

func (admin *Admin) Save(ctx context.Context) error {
	var admingroups string
	go func(lst []adm.Group) {
		for _, group := range lst {
			admingroups += strconv.FormatUint(uint64(group.Id), 10) + ","
		}
	}(admin.AdminGroup)

	var servergroups string
	go func(lst []server.Group) {
		for _, group := range lst {
			servergroups += strconv.FormatUint(uint64(group.Id), 10) + ","
		}
	}(admin.ServerGroups)

	var webgroups string
	go func(lst []web.Group) {
		for _, group := range lst {
			webgroups += strconv.FormatUint(uint64(group.Id), 10) + ","
		}
	}(admin.WebGroups)

	auths, e := json.Marshal(admin.Auths)
	if e != nil {
		return e
	}

	_, err := database.Exec(ctx, "INSERT INTO gb_admin (name, password, email, auths, flags, servers, created_at, adm_groups, web_groups, svr_groups) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", admin.Username, admin.Password, admin.Email, auths, admin.Flags, admin.Servers, admin.CreatedAt, admingroups, webgroups, servergroups)
	return err
}

func contains(lst []string, item string) bool {
	for _, i := range lst {
		if i == item {
			return true
		}
	}
	return false
}

func (admin *Admin) BuildRealFlags() string {
	//Flags with a + before character are in admin.Flags are always true
	//Flags with a - before character are in admin.Flags are always false
	//Flags with no + or - before character in admin.Flags can be negated by admin.AdminGroup.Flags
	//If a flag is negated by a group no other group can override it

	var keepFlags []string
	var negatedFlags []string
	var groupFlags []string
	for _, flag := range strings.Split(admin.Flags, ",") {
		if strings.HasPrefix(flag, "+") {
			keepFlags = append(keepFlags, strings.Replace(flag, "+", "", 1))
		} else if strings.HasPrefix(flag, "-") {
			negatedFlags = append(negatedFlags, strings.Replace(flag, "-", "", 1))
		} else {
			groupFlags = append(groupFlags, flag)
		}
	}

	sort.Slice(admin.AdminGroup, func(i, j int) bool {
		return admin.AdminGroup[i].Immunity > admin.AdminGroup[j].Immunity
	})

	for _, group := range admin.AdminGroup {
		for _, flag := range strings.Split(group.Flags, ",") {
			if strings.HasPrefix(flag, "+") && !contains(keepFlags, strings.Replace(flag, "+", "", 1)) && !contains(negatedFlags, strings.Replace(flag, "+", "", 1)) {
				keepFlags = append(keepFlags, strings.Replace(flag, "+", "", 1))
			} else if strings.HasPrefix(flag, "-") && !contains(negatedFlags, strings.Replace(flag, "-", "", 1)) && !contains(keepFlags, strings.Replace(flag, "-", "", 1)) {
				negatedFlags = append(negatedFlags, strings.Replace(flag, "-", "", 1))
			} else if !contains(groupFlags, flag) {
				groupFlags = append(groupFlags, flag)
			}
		}
	}

	var real string
	real += strings.Join(keepFlags, "")
	for _, flag := range groupFlags {
		if !contains(negatedFlags, flag) {
			real += flag
		}
	}
	return real
}
