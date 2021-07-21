package httpuserfunctions

import (
	"GoChallenge/src/frameworklayer/dependencyinjection"
	users2 "GoChallenge/src/usecaselayer/usecases/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsers(c *gin.Context) {
	var repositoryImp = dependencyinjection.GetUserRepositoryImp()
	users := users2.GetAllUsersUseCase(repositoryImp)
	c.JSON(http.StatusOK, &users)
}
