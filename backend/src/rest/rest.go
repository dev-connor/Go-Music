package rest

import (
	"backend/src/src/rest/handler"
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	r := gin.Default()

	h, _ := handler.NewHandler()

	r.GET("/products", h.GetProducts)

	r.GET("/promos", h.GetPromos)

	r.POST("/users/signin", h.SignIn)

	r.POST("/users", h.AddUser)

	r.POST("/user/:id/signout", h.SignOut)

	r.GET("/user/:id/orders", h.GetOrders)

	r.POST("/users/charge", h.Charge)

	return r.Run(address)
}
