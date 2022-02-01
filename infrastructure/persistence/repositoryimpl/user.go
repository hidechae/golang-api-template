package repositoryimpl

import (
	"github.com/hidechae/golang-api-template/core/domain/model"
	"github.com/hidechae/golang-api-template/core/domain/model/identifier"
	"github.com/hidechae/golang-api-template/core/domain/repository"
)

type UserRepositoryImpl struct {
}

var _ repository.UserRepository = UserRepositoryImpl{}

func (u UserRepositoryImpl) Find(id identifier.UserID) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}
