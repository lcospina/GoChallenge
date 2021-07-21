package http_user_functions

import (
	"GoChallenge/src/framework_layer/dependency_injection"
	"GoChallenge/src/framework_layer/utils"
	"GoChallenge/src/usecase_layer/usecases/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSeedsDB(c *gin.Context) {
	var repositoryImp = dependency_injection.GetUserRepositoryImp()
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
