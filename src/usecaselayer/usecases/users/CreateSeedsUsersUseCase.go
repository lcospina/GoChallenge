package users

import (
	"GoChallenge/src/usecaselayer/repositories"
)

func CreateSeedsUsersUseCase(userRepository repositories.UserRepository) bool {
	return userRepository.CreateSeedsUsers()
}
