package userusecases

import (
	"GoChallenge/src/domains/models"
	"GoChallenge/src/usecases/repositories"
)

func GetAllUsers(userRepository repositories.UserRepository) []models.User {
	return userRepository.GetAll()
}
