package repository

import (
	"golang-poc-postgress/domain/model"

	"github.com/jinzhu/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

type RoleRepository interface {
	FindAll(u []*model.Role) ([]*model.Role, error)
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db}
}

func (ur *roleRepository) FindAll(u []*model.Role) ([]*model.Role, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}
