package repositories

import "GoChallenge/src/domain_layer/models"

type NumbersRepository interface {
	OrderNumbers(numbers models.Numbers) models.Numbers
}
