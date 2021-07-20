package userusecases

import (
	"GoChallenge/src/usecases/repositories"
)

func CreateSeedsUsers(userRepository repositories.UserRepository) {
	userRepository.CreateSeedsUsers()
}
