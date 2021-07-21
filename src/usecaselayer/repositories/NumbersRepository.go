package repositories

import "GoChallenge/src/domainlayer/models"

type NumbersRepository interface {
	OrderNumbers(numbers models.Numbers) models.Numbers
}
