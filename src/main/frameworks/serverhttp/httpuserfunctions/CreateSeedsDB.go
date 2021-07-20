package httpuserfunctions

import (
	"GoChallenge/src/main/frameworks/dependencyinjection"
	"GoChallenge/src/main/frameworks/utils"
	"GoChallenge/src/usecases/userusecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSeedsDB(c *gin.Context) {
	var repositoryImp = dependencyinjection.GetUserRepositoryImp()
	userusecases.CreateSeedsUsersUseCase(repositoryImp)
	c.JSON(http.StatusOK, utils.USER_SEEDS_CREATE_OK)
}
