package users

import (
	"GoChallenge/src/domains/models"
	"GoChallenge/src/usecases/repositories"
)

func GetAllUsersUseCase(userRepository repositories.UserRepository) []models.User {
	return userRepository.GetAll()
}
