package numbers

import (
	"GoChallenge/src/domain_layer/models"
	"GoChallenge/src/usecase_layer/repositories"
)

func OrderNumbers(repository repositories.NumbersRepository, numbers models.Numbers) models.Numbers {
	return repository.OrderNumbers(numbers)
}
