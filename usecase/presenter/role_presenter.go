package presenter

import "golang-poc-postgress/domain/model"

type RolePresenter interface {
	ResponseRoles(u []*model.Role) []*model.Role
}
