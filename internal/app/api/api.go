package api

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *gin.Engine {
	router := gin.New()

	walletGroup := router.Group("/wallet")
	{
		walletGroup.GET("find/all", nil)
		walletGroup.GET("find/email/:email", nil)
		walletGroup.GET("find/id/:id", nil)
		walletGroup.GET("find/phone/:phone", nil)
		walletGroup.GET("history", nil)

		walletGroup.POST("create", nil)
		walletGroup.POST("deposit", nil)
		walletGroup.POST("withdraw", nil)
		walletGroup.POST("transfer", nil)

		walletGroup.PATCH("update/:id", nil)

		walletGroup.DELETE("delete/:id", nil)
	}

	return router
}
