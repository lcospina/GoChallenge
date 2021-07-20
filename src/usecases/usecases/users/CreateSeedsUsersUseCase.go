package users

import (
	"GoChallenge/src/usecases/repositories"
)

func CreateSeedsUsersUseCase(userRepository repositories.UserRepository) bool {
	return userRepository.CreateSeedsUsers()
}
