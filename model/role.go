package model

import "errors"

type Role string

type Policy string

type ServerPolicy string

type ClientPolicy string

const (
	ADMIN Role = "ADMIN"
)

func GetRoles() []Role {
	return []Role{
		ADMIN,
	}
}

func GetRole(roleStr string) (Role, error) {
	for _, r := range GetRoles() {
		if string(r) == roleStr {
			return r, nil
		}
	}
	return "", errors.New("No role exist for this val")
}
