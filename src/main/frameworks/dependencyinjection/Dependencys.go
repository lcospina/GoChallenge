package dependencyinjection

import (
	local "GoChallenge/src/main/frameworks/repositoriesimps/local"
	mysql "GoChallenge/src/main/frameworks/repositoriesimps/mysql"
	"GoChallenge/src/usecases/repositories"
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
