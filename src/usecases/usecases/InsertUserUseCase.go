package usecases

import (
	"GoChallenge/src/domains/models"
	"GoChallenge/src/usecases/repositories"
)

func InsertUserUseCase(userRepository repositories.UserRepository, user models.User) bool {
	return userRepository.Insert(user)
}
