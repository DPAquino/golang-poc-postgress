package registry

import (
	"golang-poc-postgress/interface/controller"
	ip "golang-poc-postgress/interface/presenter"
	ir "golang-poc-postgress/interface/repository"
	"golang-poc-postgress/usecase/interactor"
	up "golang-poc-postgress/usecase/presenter"
	ur "golang-poc-postgress/usecase/repository"
)

func (r *registry) NewRoleController() controller.RoleController {
	return controller.NewRoleController(r.NewRoleInteractor())
}

func (r *registry) NewRoleInteractor() interactor.RoleInteractor {
	return interactor.NewRoleInteractor(r.NewRoleRepository(), r.NewRolePresenter())
}

func (r *registry) NewRoleRepository() ur.RoleRepository {
	return ir.NewRoleRepository(r.db)
}

func (r *registry) NewRolePresenter() up.RolePresenter {
	return ip.NewRolePresenter()
}
