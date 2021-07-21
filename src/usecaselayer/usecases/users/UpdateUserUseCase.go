package users

import (
	"GoChallenge/src/domainlayer/models"
	"GoChallenge/src/usecaselayer/repositories"
)

func UpdateUserUseCase(userRepository repositories.UserRepository, user models.User) bool {
	return userRepository.Update(user)
}
