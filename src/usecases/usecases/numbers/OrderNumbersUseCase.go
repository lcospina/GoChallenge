package numbers

import (
	"GoChallenge/src/domains/models"
	"GoChallenge/src/usecases/repositories"
)

func OrderNumbers(repository repositories.NumbersRepository, numbers models.Numbers) models.Numbers {
	return repository.OrderNumbers(numbers)
}
