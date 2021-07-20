package main

import (
	"GoChallenge/src/main/frameworks/serverhttp"
)

func main() {

	//PARA POBLAR LA BD, Esto es solo para prueba...
	//var repositoryImp = dependencyinjection.GetUserRepositoryImp()
	//userusecases.CreateSeedsUsers(repositoryImp)

	serverhttp.InitServer("9090")

}
