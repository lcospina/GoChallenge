package userfunctions

import (
	"GoChallenge/src/main/frameworks/repositoriesimplements"
	"GoChallenge/src/usecases/userusecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsers(c *gin.Context) {
	var repositoryImp = new(repositoriesimplements.UserRepositoryMysql)
	users := userusecases.GetAllUsers(repositoryImp)
	c.JSON(http.StatusOK, &users)
}
