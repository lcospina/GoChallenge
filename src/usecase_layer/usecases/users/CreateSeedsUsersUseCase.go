package users

import (
	"GoChallenge/src/usecase_layer/repositories"
)

func CreateSeedsUsersUseCase(userRepository repositories.UserRepository) bool {
	return userRepository.CreateSeedsUsers()
}
