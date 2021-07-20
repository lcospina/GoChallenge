package userusecases

import (
	"GoChallenge/src/domains/models"
	"GoChallenge/src/usecases/repositories"
)

func DeleteUserUseCase(userRepository repositories.UserRepository, user models.User) bool {
	return userRepository.Delete(user)
}
