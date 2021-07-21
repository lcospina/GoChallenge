package users

import (
	"GoChallenge/src/domainlayer/models"
	"GoChallenge/src/usecaselayer/repositories"
)

func InsertUserUseCase(userRepository repositories.UserRepository, user models.User) bool {
	return userRepository.Insert(user)
}
