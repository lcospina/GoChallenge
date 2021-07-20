package httpuserfunctions

import (
	"GoChallenge/src/main/frameworks/dependencyinjection"
	users2 "GoChallenge/src/usecases/usecases/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsers(c *gin.Context) {
	var repositoryImp = dependencyinjection.GetUserRepositoryImp()
	users := users2.GetAllUsersUseCase(repositoryImp)
	c.JSON(http.StatusOK, &users)
}
