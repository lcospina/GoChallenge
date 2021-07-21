package users

import (
	"GoChallenge/src/domainlayer/models"
	"GoChallenge/src/usecaselayer/repositories"
)

func GetAllUsersUseCase(userRepository repositories.UserRepository) []models.User {
	return userRepository.GetAll()
}
