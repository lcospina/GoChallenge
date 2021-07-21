package httpuserfunctions

import (
	"GoChallenge/src/domainlayer/models"
	"GoChallenge/src/frameworklayer/dependencyinjection"
	"GoChallenge/src/frameworklayer/utils"
	"GoChallenge/src/usecaselayer/usecases/users"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusOK, utils.USER_IVALID)
	} else {
		var repositoryImp = dependencyinjection.GetUserRepositoryImp()
		resp := users.DeleteUserUseCase(repositoryImp, models.User{ID: idInt})
		if resp {
			c.JSON(http.StatusOK, gin.H{
				utils.RESPONSE: utils.OK,
				utils.DATA:     utils.USER_DELETE_OK,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				utils.RESPONSE: utils.ERROR,
				utils.DATA:     utils.USER_DELETE_NOOK,
			})
		}
	}
}
