package httpuserfunctions

import (
	"GoChallenge/src/frameworklayer/dependencyinjection"
	"GoChallenge/src/frameworklayer/utils"
	"GoChallenge/src/usecaselayer/usecases/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSeedsDB(c *gin.Context) {
	var repositoryImp = dependencyinjection.GetUserRepositoryImp()
	var resp = users.CreateSeedsUsersUseCase(repositoryImp)
	if resp {
		c.JSON(http.StatusOK, gin.H{
			utils.RESPONSE: utils.OK,
			utils.DATA:     utils.USER_SEEDS_CREATE_OK,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			utils.RESPONSE: utils.ERROR,
			utils.DATA:     utils.USER_SEEDS_CREATE_NOOK,
		})
	}
}
