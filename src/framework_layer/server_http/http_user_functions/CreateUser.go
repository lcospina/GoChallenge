package http_user_functions

import (
	"GoChallenge/src/domain_layer/models"
	"GoChallenge/src/framework_layer/dependency_injection"
	"GoChallenge/src/framework_layer/utils"
	"GoChallenge/src/usecase_layer/usecases/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return
	}
	var repositoryImp = dependency_injection.GetUserRepositoryImp()
	resp := users.InsertUserUseCase(repositoryImp, user)
	if resp {
		c.JSON(http.StatusOK, gin.H{
			utils.RESPONSE: utils.OK,
			utils.DATA:     utils.USER_CREATE_OK,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			utils.RESPONSE: utils.ERROR,
			utils.DATA:     utils.USER_CREATE_NOOK,
		})
	}

}
