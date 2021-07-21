package dependency_injection

import (
	"GoChallenge/src/framework_layer/repositories_impl/local"
	"GoChallenge/src/framework_layer/repositories_impl/mysql"
	"GoChallenge/src/usecase_layer/repositories"
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
