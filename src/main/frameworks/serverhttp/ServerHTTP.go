package serverhttp

import (
	"GoChallenge/src/main/frameworks/serverhttp/numbersfunctions"
	"GoChallenge/src/main/frameworks/serverhttp/userfunctions"
	"github.com/gin-gonic/gin"
)

func InitServer(port string) {
	router := gin.Default()
	router.GET(ROUTE_GET_ALL_USERS, userfunctions.GetAllUsers)
	router.GET(ROUTE_ORDER_NUMBERS, numbersfunctions.OrderNumbers)
	router.Run(":" + port)
}
