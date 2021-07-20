package httpuserfunctions

import (
	"GoChallenge/src/domains/models"
	"GoChallenge/src/main/frameworks/dependencyinjection"
	"GoChallenge/src/main/frameworks/utils"
	"GoChallenge/src/usecases/usecases/users"
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
