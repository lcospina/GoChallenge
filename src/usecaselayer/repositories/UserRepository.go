package repositories

import (
	"GoChallenge/src/domainlayer/models"
)

type UserRepository interface {
	Insert(user models.User) bool

	Update(user models.User) bool

	Delete(user models.User) bool

	FindByIdUserUseCase(user models.User) models.User

	GetAll() []models.User

	CreateSeedsUsers() bool
}
