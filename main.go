package main

/*
import "GoChallenge/src/main/frameworks/serverhttp"

func main() {
	serverhttp.InitServer("9090")
}
*/

import (
	"GoChallenge/src/main/frameworks/serverhttp"
	"GoChallenge/src/main/frameworks/serverhttp/httpnumbersfunctions"
	"GoChallenge/src/main/frameworks/serverhttp/httpuserfunctions"
	"GoChallenge/src/main/frameworks/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Use(middleware)
	router.GET(serverhttp.ROUTE_GET_ALL_USERS, httpuserfunctions.GetAllUsers)
	router.DELETE(serverhttp.ROUTE_DELETE_USERS, httpuserfunctions.DeleteUser)
	router.POST(serverhttp.ROUTE_CREATE_USERS, httpuserfunctions.CreateUser)
	router.GET(serverhttp.ROUTE_CREATE_SEEDSDB, httpuserfunctions.CreateSeedsDB)
	router.GET(serverhttp.ROUTE_FIND_BY_ID_USERS, httpuserfunctions.FindUserById)
	router.POST(serverhttp.ROUTE_ORDER_NUMBERS, httpnumbersfunctions.OrderNumbers)
	err := router.Run(":80")
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
