package users

import (
	"GoChallenge/src/domain_layer/models"
	"GoChallenge/src/usecase_layer/repositories"
)

func UpdateUserUseCase(userRepository repositories.UserRepository, user models.User) bool {
	return userRepository.Update(user)
}
