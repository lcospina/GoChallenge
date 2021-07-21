package server_http

import (
	"GoChallenge/src/framework_layer/server_http/http_numbers_functions"
	"GoChallenge/src/framework_layer/server_http/http_user_functions"
	"GoChallenge/src/framework_layer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitServer(port string) {
	router := gin.Default()
	router.Use(middleware)
	router.GET(ROUTE_GET_ALL_USERS, http_user_functions.GetAllUsers)
	router.DELETE(ROUTE_DELETE_USERS, http_user_functions.DeleteUser)
	router.POST(ROUTE_CREATE_USERS, http_user_functions.CreateUser)
	router.GET(ROUTE_CREATE_SEEDSDB, http_user_functions.CreateSeedsDB)
	router.GET(ROUTE_FIND_BY_ID_USERS, http_user_functions.FindUserById)
	router.POST(ROUTE_ORDER_NUMBERS, http_numbers_functions.OrderNumbers)
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
