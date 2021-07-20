package httpuserfunctions

import (
	"GoChallenge/src/main/frameworks/dependencyinjection"
	"GoChallenge/src/usecases/userusecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsers(c *gin.Context) {
	var repositoryImp = dependencyinjection.GetUserRepositoryImp()
	users := userusecases.GetAllUsers(repositoryImp)
	c.JSON(http.StatusOK, &users)
}
