package http_user_functions

import (
	"GoChallenge/src/framework_layer/dependency_injection"
	users2 "GoChallenge/src/usecase_layer/usecases/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsers(c *gin.Context) {
	var repositoryImp = dependency_injection.GetUserRepositoryImp()
	users := users2.GetAllUsersUseCase(repositoryImp)
	c.JSON(http.StatusOK, &users)
}
