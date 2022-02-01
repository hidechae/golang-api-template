package model

import (
	"github.com/hidechae/golang-api-template/core/domain/model/identifier"
)

type User struct {
	ID   identifier.UserID
	Name string
}
