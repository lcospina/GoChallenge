package users

import (
	"GoChallenge/src/domainlayer/models"
	"GoChallenge/src/usecaselayer/repositories"
)

func DeleteUserUseCase(userRepository repositories.UserRepository, user models.User) bool {
	return userRepository.Delete(user)
}
