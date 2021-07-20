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

func FindUserById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusOK, utils.USER_IVALID)
	} else {
		var repositoryImp = dependencyinjection.GetUserRepositoryImp()
		userTemp := users.FindByIdUserUseCase(repositoryImp, models.User{ID: idInt})

		if userTemp.ID != 0 {
			c.JSON(http.StatusOK, gin.H{
				utils.RESPONSE: utils.OK,
				utils.DATA:     userTemp,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				utils.RESPONSE: utils.ERROR,
				utils.DATA:     utils.USER_FIND_NOOK,
			})
		}
	}
}
