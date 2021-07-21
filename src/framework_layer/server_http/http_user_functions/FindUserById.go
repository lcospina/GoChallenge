package http_user_functions

import (
	"GoChallenge/src/domain_layer/models"
	"GoChallenge/src/framework_layer/dependency_injection"
	"GoChallenge/src/framework_layer/utils"
	"GoChallenge/src/usecase_layer/usecases/users"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindUserById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			utils.RESPONSE: utils.ERROR,
			utils.DATA:     utils.USER_IVALID,
		})
	} else {
		var repositoryImp = dependency_injection.GetUserRepositoryImp()
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
