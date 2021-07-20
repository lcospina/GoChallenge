package userusecases

import (
	"GoChallenge/src/domains/models"
	"GoChallenge/src/usecases/repositories"
)

func UpdateUserUseCase(userRepository repositories.UserRepository, user models.User) bool {
	return userRepository.Update(user)
}
