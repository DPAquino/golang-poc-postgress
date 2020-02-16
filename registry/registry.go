package registry

import (
	"golang-poc-postgress/interface/controller"

	"github.com/jinzhu/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	UserController() controller.UserController
	RoleController() controller.RoleController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) UserController() controller.UserController {
	return r.NewUserController()
}

func (r *registry) RoleController() controller.RoleController {
	return r.NewRoleController()
}
