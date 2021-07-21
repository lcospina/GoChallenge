package httpuserfunctions

import (
	"GoChallenge/src/domainlayer/models"
	"GoChallenge/src/frameworklayer/dependencyinjection"
	"GoChallenge/src/frameworklayer/utils"
	"GoChallenge/src/usecaselayer/usecases/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return
	}
	var repositoryImp = dependencyinjection.GetUserRepositoryImp()
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
