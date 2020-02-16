package repository

import "golang-poc-postgress/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
