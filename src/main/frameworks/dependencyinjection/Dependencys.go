package dependencyinjection

import (
	"GoChallenge/src/main/frameworks/repositoriesimps/mysql"
	"GoChallenge/src/usecases/repositories"
)

var userRepositoryImp repositories.UserRepository

func GetUserRepositoryImp() repositories.UserRepository {
	if userRepositoryImp == nil {
		userRepositoryImp = new(mysql.UserRepositoryMysql)
	}
	return userRepositoryImp
}
