package users

import (
	"GoChallenge/src/domain_layer/models"
	"GoChallenge/src/usecase_layer/repositories"
)

func FindByIdUserUseCase(userRepository repositories.UserRepository, user models.User) models.User {
	return userRepository.FindByIdUserUseCase(user)
}
