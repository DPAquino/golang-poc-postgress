package controller

import (
	"golang-poc-postgress/domain/model"
	"golang-poc-postgress/usecase/interactor"
	"net/http"
)

type roleController struct {
	roleInteractor interactor.RoleInteractor
}

type RoleController interface {
	GetRoles(c Context) error
}

func NewRoleController(us interactor.RoleInteractor) RoleController {
	return &roleController{us}
}

func (uc *roleController) GetRoles(c Context) error {
	var u []*model.Role

	u, err := uc.roleInteractor.Get(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
