package repositories

import "GoChallenge/src/domains/models"

type NumbersRepository interface {
	OrderNumbers(numbers models.Numbers) models.Numbers
}
