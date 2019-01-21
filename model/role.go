package model

type Role struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}

type Policy string

type ServerPolicy string

type ClientPolicy string

const (
	ADMIN       = "ADMIN"
	SUPER_ADMIN = "SUPER_ADMIN"
)

var roles map[string]*Role

func InitRoles() map[string]*Role {
	roles = make(map[string]*Role)
	roles[SUPER_ADMIN] = &Role{
		ID:          "SUPER_ADMIN",
		Description: "Super admin role who can do anything",
		Permissions: []string{
			PERMISSION_SUDO_USER.ID,
		},
	}
	return roles
}

func GetRoles() []string {
	return []string{
		ADMIN,
		SUPER_ADMIN,
	}
}

func GetRole(roleStr string) *Role {
	return roles[roleStr]
}
