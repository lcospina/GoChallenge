package usecases

import (
	"GoChallenge/src/domains/models"
	"GoChallenge/src/usecases/repositories"
)

func SelectUserFindIdUserUseCase(userRepository repositories.UserRepository, user models.User) models.User {
	return userRepository.SelectFindId(user)
}
