package users

import (
	"GoChallenge/src/domainlayer/models"
	"GoChallenge/src/usecaselayer/repositories"
)

func FindByIdUserUseCase(userRepository repositories.UserRepository, user models.User) models.User {
	return userRepository.FindByIdUserUseCase(user)
}
