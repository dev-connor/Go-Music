# Go-Music

React 와 Go 를 이용해 악기 판매사이트를 제작했습니다. 

![Home](https://user-images.githubusercontent.com/70655507/216764186-e9b22a9c-61af-416c-9334-4220803dc82c.png)

Stack: React, Go, Gin, GORM

개발인원: 개인

개발기간
- 2022년 10월 7일 ~ 2022년 10월 15일 (약 일주일)

> 개발과정은 여기 블로그에서 볼 수 있습니다. <br>
https://devconnor.tistory.com/113

<br><br>

Frontend
- 목록 뿌리기 
- Proxy 로 포워딩 

Backend
- 라우터 정의
- 커스텀 미들웨어 정의
- 자체 서명 디지털 인증서 발급받기
- 패스워드 해싱 (Password hashing) 
- 신용카드 결제요청하기 (스프라이트 api) 

---

<br><br><br>

1. 네비게이션 목록

![Home](https://user-images.githubusercontent.com/70655507/216764220-db33ef2c-383d-425a-b50d-a9786b257f1b.png)

<img width="1427" alt="products" src="https://user-images.githubusercontent.com/70655507/216764233-7d6812ca-8fb2-44d4-9f81-8ce6597674b9.png">

> DB 에 있는 악기 정보들을 조회하는 쿼리를

<img width="964" alt="쿼리 로그" src="https://user-images.githubusercontent.com/70655507/216764238-30a52085-b17e-4d3c-a64d-e9cab2fc3f49.png">

> 로그를 통해 확인할 수 있습니다.

![Promotions](https://user-images.githubusercontent.com/70655507/216764224-58dff736-aa53-4e74-b742-c2aa01ec9e00.png)

> Promotions 탭에서는 할인이 적용된 상품들의 목록을 볼 수 있습니다. 

2. 신용카드 결제요청하기

https://user-images.githubusercontent.com/70655507/216765362-3582628b-985e-4925-bb25-54d7c096d394.mp4

> 유효한 카드인지 검증하는 모습을 볼 수 있다.

3. 라우팅 정의
```go
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
```

4. 테이블

![image](https://user-images.githubusercontent.com/70655507/216621139-5240c27b-947a-471b-9489-1a7347db8e6d.png)

5. 커스텀 미들웨어 작성

```go
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
```

```go
func RunAPIWithHandler(address string, h HandlerInterface) error {
    r := gin.Default()
    r.Use(MyCustomerLogger())
```

> Gin 의 기본 미들웨어는 유지하고 MyCustomLogger() 라는 새로운 커스텀 미들웨어를 추가한다.

6. 자체 서명 디지털 인증서 발급

Go 의 기본 라이브러리를 사용하여 인증서를 발급해보자.

![image](https://user-images.githubusercontent.com/70655507/216622123-c348640d-1754-4698-ac5c-4f24bf293737.png)


![image](https://user-images.githubusercontent.com/70655507/216622144-3ea418e2-4012-4eba-86a1-7b5ffb55c068.png)

> 두 파일이 생긴 것을 볼 수 있다.

7. 패스워드 해싱 (Password hashing)

![Sign in](https://user-images.githubusercontent.com/70655507/216764251-92bcd704-a034-4fd5-b0f3-83e99014a26e.png)

위 함수는 bcrypt 패키지를 사용한다. 이 패키지는 패스워드 해싱에 주로 사용된다.

bcrypt 는 1990 년대에 설계된 유명한 해싱 기법이다.

OpenBSD 운영체제의 기본 패스워드 해싱 기법이며, 많은 프로그래밍 언어에서 지원한다.

bcrypt 패키지는 패스워드 해시와 일반 문자열을 비교하는 메서드도 제공한다.

```go
func hashPassword(s *string) error {
    if s == nil {
        return errors.New("Reference provided for hashing password is nil")
    }
    // bcrypt 패키지에서 사용할 수 있게 패스워드 문자열을 바이트 슬라이스로 변환한다.
    sBytes := []byte(*s)
    // GenerateFromPassword() 메서드는 패스워드 해시를 반환한다.
    hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    // 패스워드 문자열을 해시 값으로 바꾼다.
    *s = string(hashedBytes[:])
    return nil
}
```


