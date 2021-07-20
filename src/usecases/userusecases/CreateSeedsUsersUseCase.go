package userusecases

import (
	"GoChallenge/src/usecases/repositories"
)

func CreateSeedsUsersUseCase(userRepository repositories.UserRepository) {
	userRepository.CreateSeedsUsers()
}
