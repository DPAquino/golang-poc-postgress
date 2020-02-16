package presenter

import "golang-poc-postgress/domain/model"

type rolePresenter struct {
}

type RolePresenter interface {
	ResponseRoles(us []*model.Role) []*model.Role
}

func NewRolePresenter() RolePresenter {
	return &rolePresenter{}
}

func (up *rolePresenter) ResponseRoles(us []*model.Role) []*model.Role {
	for _, u := range us {
		u.RoleName = u.RoleName
	}
	return us
}
