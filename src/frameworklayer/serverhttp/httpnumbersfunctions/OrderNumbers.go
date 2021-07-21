package httpnumbersfunctions

import (
	"GoChallenge/src/domainlayer/models"
	"GoChallenge/src/frameworklayer/dependencyinjection"
	"GoChallenge/src/frameworklayer/utils"
	"GoChallenge/src/usecaselayer/usecases/numbers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func OrderNumbers(c *gin.Context) {
	var dataParsed models.Numbers
	err := c.Bind(&dataParsed)
	if err != nil {
		return
	}

	var repositoryImp = dependencyinjection.GetNumberRepositoryImp()
	resp := numbers.OrderNumbers(repositoryImp, dataParsed)

	c.JSON(http.StatusOK, gin.H{
		utils.RESPONSE: utils.OK,
		utils.DATA:     resp,
	})
}
