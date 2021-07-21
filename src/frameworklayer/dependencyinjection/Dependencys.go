package dependencyinjection

import (
	"GoChallenge/src/frameworklayer/repositoriesimps/local"
	"GoChallenge/src/frameworklayer/repositoriesimps/mysql"
	"GoChallenge/src/usecaselayer/repositories"
)

var userRepositoryImp repositories.UserRepository
var numberRepositoryImp repositories.NumbersRepository

func GetUserRepositoryImp() repositories.UserRepository {
	if userRepositoryImp == nil {
		userRepositoryImp = new(mysql.UserRepositoryMysql)
	}
	return userRepositoryImp
}

func GetNumberRepositoryImp() repositories.NumbersRepository {
	if numberRepositoryImp == nil {
		numberRepositoryImp = new(local.NumbersRepositoryLocal)
	}
	return numberRepositoryImp
}
