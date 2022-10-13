package rest

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func RunAPIWithHandler(address string, h HandlerInterface) error {
	r := gin.Default()
	r.Use(MyCustomerLogger())
	r.Use(static.ServeRoot("/", "../public/build"))

	h, _ = NewHandler()

	r.GET("/products", h.GetProducts)

	r.GET("/promos", h.GetPromos)

	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.GET("/:id/orders", h.GetOrders)
	}
	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/charge", h.Charge)
		usersGroup.POST("/signin", h.SignIn)
		usersGroup.POST("", h.AddUser)
	}

	return r.Run(address)
}

func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}

func MyCustomerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 요청을 처리하기 전에 실행할 코드
		// 예제 변수 설정
		c.Set("v", "123")
		// c.Get("v) 를 요청하면 변수 값을 확인할 수 있다.

		// 요청 처리 로직 실행
		c.Next()

		// 이 코드는 핸들러 실행이 끝나면 실행된다.

		// 응답코드 확인
		status := c.Writer.Status()
		// status 를 활용하는 코드 추가
	}
}

func MyCustomerLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("**********************")
		c.Next()
		fmt.Println("**********************")
	}
}
