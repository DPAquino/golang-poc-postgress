package interactor

import (
	"golang-poc-postgress/domain/model"
	"golang-poc-postgress/usecase/presenter"
	"golang-poc-postgress/usecase/repository"
)

type roleInteractor struct {
	RoleRepository repository.RoleRepository
	RolePresenter  presenter.RolePresenter
}

type RoleInteractor interface {
	Get(u []*model.Role) ([]*model.Role, error)
}

func NewRoleInteractor(r repository.RoleRepository, p presenter.RolePresenter) RoleInteractor {
	return &roleInteractor{r, p}
}

func (us *roleInteractor) Get(u []*model.Role) ([]*model.Role, error) {
	u, err := us.RoleRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.RolePresenter.ResponseRoles(u), nil
}
