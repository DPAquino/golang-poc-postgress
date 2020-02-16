package repository

import "golang-poc-postgress/domain/model"

type RoleRepository interface {
	FindAll(u []*model.Role) ([]*model.Role, error)
}
