package presenter

import "golang-poc-postgress/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
