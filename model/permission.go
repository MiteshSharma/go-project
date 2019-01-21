package model

type Permission struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

var allPermissions []*Permission
var PERMISSION_SUDO_USER *Permission

func InitPermissions() {
	PERMISSION_SUDO_USER = &Permission{
		"all_sudo_user_tasks",
		"Permission to do all super user tasks",
	}

	allPermissions = []*Permission{
		PERMISSION_SUDO_USER,
	}
}
