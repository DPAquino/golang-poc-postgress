package registry

import (
	"golang-poc-postgress/interface/controller"
	ip "golang-poc-postgress/interface/presenter"
	ir "golang-poc-postgress/interface/repository"
	"golang-poc-postgress/usecase/interactor"
	up "golang-poc-postgress/usecase/presenter"
	ur "golang-poc-postgress/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
