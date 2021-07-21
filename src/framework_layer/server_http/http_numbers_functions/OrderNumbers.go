package http_numbers_functions

import (
	"GoChallenge/src/domain_layer/models"
	"GoChallenge/src/framework_layer/dependency_injection"
	"GoChallenge/src/framework_layer/utils"
	"GoChallenge/src/usecase_layer/usecases/numbers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func OrderNumbers(c *gin.Context) {
	var dataParsed models.Numbers
	err := c.Bind(&dataParsed)
	if err != nil {
		return
	}

	var repositoryImp = dependency_injection.GetNumberRepositoryImp()
	resp := numbers.OrderNumbers(repositoryImp, dataParsed)

	c.JSON(http.StatusOK, gin.H{
		utils.RESPONSE: utils.OK,
		utils.DATA:     resp,
	})
}
