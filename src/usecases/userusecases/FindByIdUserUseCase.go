package userusecases

import (
	"GoChallenge/src/domains/models"
	"GoChallenge/src/usecases/repositories"
)

func FindByIdUserUseCase(userRepository repositories.UserRepository, user models.User) models.User {
	return userRepository.FindByIdUserUseCase(user)
}
