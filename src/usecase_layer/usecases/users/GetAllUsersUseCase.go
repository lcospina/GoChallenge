package users

import (
	"GoChallenge/src/domain_layer/models"
	"GoChallenge/src/usecase_layer/repositories"
)

func GetAllUsersUseCase(userRepository repositories.UserRepository) []models.User {
	return userRepository.GetAll()
}
