package main

import "GoChallenge/src/main/frameworks/serverhttp"

func main() {

	serverhttp.InitServer("9090")
	/*
		user := *new(models.User)
		user.FirstName = "Luis"
		user.LastName = "Ospina"
		user.NumberPhone = "3173040456"
		user.Email = "lcospina9@gmail.com"
		user.CreatedAt = time.Time{}

		role := *new(models.Role)
		role.ID = 1
		role.Description = "Administrador"

		user.Role = role

		var repositoryImp = new(repositoriesimplements.UserRepositoryMysql)
		usecases.InsertUserUseCase(repositoryImp, user)
	*/
}
