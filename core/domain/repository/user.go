package repository

import (
	"github.com/hidechae/golang-api-template/core/domain/model"
	"github.com/hidechae/golang-api-template/core/domain/model/identifier"
)

type UserRepository interface {
	Find(identifier.UserID) (*model.User, error)
}
