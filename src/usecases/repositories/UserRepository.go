package repositories

import (
	"GoChallenge/src/domains/models"
)

type UserRepository interface {
	Insert(user models.User) bool

	Update(user models.User) bool

	Delete(user models.User) bool

	SelectFindId(user models.User) models.User

	GetAll() []models.User

	CreateSeedsUsers()
}
