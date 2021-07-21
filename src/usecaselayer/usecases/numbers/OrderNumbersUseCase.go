package numbers

import (
	"GoChallenge/src/domainlayer/models"
	"GoChallenge/src/usecaselayer/repositories"
)

func OrderNumbers(repository repositories.NumbersRepository, numbers models.Numbers) models.Numbers {
	return repository.OrderNumbers(numbers)
}
