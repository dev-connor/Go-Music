package rest

import "github.com/gin-gonic/gin"

func RunAPI(address string) error {
	r := gin.Default()
	r.GET("/relativepath/to/url", func(c *gin.Context) {
		// 로직 구현
	})

	// 상품 목록
	r.GET("/products", func(c *gin.Context) {
		// 클라이언트에게 상품 목록 반환

	})

	// 프로모션 목록
	r.GET("/promos", func(c *gin.Context) {

	})

	// 사용자 로그인 POST 요청
	r.POST("/users/signin", func(c *gin.Context) {
		// 사용자 로그인

	})

	// 사용자 추가 POST 요청
	r.POST("/users", func(c *gin.Context) {
		// 사용자 추가

	})

	// 사용자 로그아웃 POST 요청
	/*
		아래 경로는 사용자 ID 를 포함한다. ID 는 사용자마다 고유한 값이기 때문에
		와일드카드 (*) 를 사용한다. ':id' 는 변수 id 를 의미한다.
	*/
	r.POST("/user/:id/signout", func(c *gin.Context) {
		// 해당 ID 의 사용자 로그아웃

	})

	// 구매 목록 조회
	r.GET("/user/:id/orders", func(c *gin.Context) {
		// 해당 ID 의 사용자의 주문내역 조회

	})

	// 결제 POST 요청
	r.POST("/users/charge", func(c *gin.Context) {
		// 신용카드 결제 처리

	})

}
