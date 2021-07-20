package httpuserfunctions

import (
	"GoChallenge/src/domains/models"
	"GoChallenge/src/main/frameworks/dependencyinjection"
	"GoChallenge/src/main/frameworks/utils"
	"GoChallenge/src/usecases/userusecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	var repositoryImp = dependencyinjection.GetUserRepositoryImp()
	resp := userusecases.InsertUserUseCase(repositoryImp, user)
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