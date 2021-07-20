package serverhttp

import (
	"GoChallenge/src/main/frameworks/serverhttp/httpnumbersfunctions"
	"GoChallenge/src/main/frameworks/serverhttp/httpuserfunctions"
	"GoChallenge/src/main/frameworks/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitServer(port string) {
	router := gin.Default()
	router.Use(middleware)
	router.GET(ROUTE_GET_ALL_USERS, httpuserfunctions.GetAllUsers)
	router.DELETE(ROUTE_DELETE_USERS, httpuserfunctions.DeleteUser)
	router.POST(ROUTE_CREATE_USERS, httpuserfunctions.CreateUser)
	router.GET(ROUTE_CREATE_SEEDSDB, httpuserfunctions.CreateSeedsDB)
	router.GET(ROUTE_FIND_BY_ID_USERS, httpuserfunctions.FindUserById)
	router.POST(ROUTE_ORDER_NUMBERS, httpnumbersfunctions.OrderNumbers)
	err := router.Run(":" + port)
	if err != nil {
		println(utils.SERVER_ERROR_INIT)
		return
	}
}

func middleware(c *gin.Context) {
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": utils.SERVER_ERROR_INTERNAL,
			})
		}
	}(c)
	c.Next()
}
