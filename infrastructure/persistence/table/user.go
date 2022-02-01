package table

import (
	"github.com/hidechae/golang-api-template/core/domain/model"
	"github.com/hidechae/golang-api-template/core/domain/model/identifier"
)

type UserTable struct {
	ID   uint64
	Name string
}

func NewUserTable(user model.User) UserTable {
	return UserTable{
		ID:   user.ID.Int(),
		Name: user.Name,
	}
}

func (t *UserTable) ToModel() model.User {
	return model.User{
		ID:   identifier.NewUserID(t.ID),
		Name: t.Name,
	}
}
