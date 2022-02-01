package usecase

import "github.com/hidechae/golang-api-template/core/domain/repository"

type UserUseCase struct {
	userRepository repository.UserRepository
}

func (u *UserUseCase) Register() {

}
