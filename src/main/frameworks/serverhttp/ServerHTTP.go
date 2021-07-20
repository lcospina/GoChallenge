package serverhttp

import (
	"GoChallenge/src/main/frameworks/serverhttp/httpnumbersfunctions"
	"GoChallenge/src/main/frameworks/serverhttp/httpuserfunctions"
	"github.com/gin-gonic/gin"
)

func InitServer(port string) {
	router := gin.Default()
	router.GET(ROUTE_GET_ALL_USERS, httpuserfunctions.GetAllUsers)
	router.DELETE(ROUTE_DELETE_USERS, httpuserfunctions.DeleteUser)
	router.POST(ROUTE_CREATE_USERS, httpuserfunctions.CreateUser)
	router.GET(ROUTE_CREATE_SEEDSDB, httpuserfunctions.CreateSeedsDB)
	router.GET(ROUTE_FIND_BY_ID_USERS, httpuserfunctions.FindUserById)
	router.GET(ROUTE_ORDER_NUMBERS, httpnumbersfunctions.OrderNumbers)
	router.Run(":" + port)
}
