package serverhttp

import (
	"GoChallenge/src/main/frameworks/serverhttp/httpnumbersfunctions"
	"GoChallenge/src/main/frameworks/serverhttp/httpuserfunctions"
	"github.com/gin-gonic/gin"
)

func InitServer(port string) {
	router := gin.Default()
	router.GET(ROUTE_GET_ALL_USERS, httpuserfunctions.GetAllUsers)
	router.GET(ROUTE_DELETE_USERS, httpuserfunctions.DeleteUser)
	router.GET(ROUTE_ORDER_NUMBERS, httpnumbersfunctions.OrderNumbers)
	router.POST(ROUTE_CREATE_USERS, httpuserfunctions.CreateUser)
	router.GET(ROUTE_CREATE_SEEDSDB, httpuserfunctions.CreateSeedsDB)
	router.Run(":" + port)
}
