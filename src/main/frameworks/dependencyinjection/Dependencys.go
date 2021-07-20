package dependencyinjection

import (
	"GoChallenge/src/main/frameworks/repositoriesimps"
	"GoChallenge/src/usecases/repositories"
)

var userRepositoryImp repositories.UserRepository

func GetUserRepositoryImp() repositories.UserRepository {
	if userRepositoryImp == nil {
		userRepositoryImp = new(repositoriesimps.UserRepositoryMysql)
	}
	return userRepositoryImp
}
