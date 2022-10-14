package rest

import (
	"backend/src/src/dblayer"
	"backend/src/src/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHandler_GetProducts(t *testing.T) {
	// 로그가 너무 많이 쌓이지 않게 테스트 모드로 전환
	gin.SetMode(gin.TestMode)
	mockdbLayer := dblayer.NewMockDBLayerWithData()
	h := NewHandlerWithDB(mockdbLayer)
	const productsURL string = "/products"

	type errMSG struct {
		Error string `json:"error"`
	}

	tests := []struct {
		name             string
		inErr            error
		outStatusCode    int
		expectedRespBody interface{}
	}{
		{
			"getproductsnoerrors",
			nil,
			http.StatusOK,
			mockdbLayer.GetMockProductData(),
		},
		{
			"getproductswitherror",
			errors.New("get products error"),
			http.StatusInternalServerError,
			errMSG{Error: "get products error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 서브테스트 실행
			mockdbLayer.SetError(tt.inErr)
			// 테스트 요청 생성
			req := httptest.NewRequest(http.MethodGet, productsURL, nil)

			// http response recorder 생성
			w := httptest.NewRecorder()
			_, engine := gin.CreateTestContext(w)

			// get 요청 설정
			engine.GET(productsURL, h.GetProducts)
			engine.ServeHTTP(w, req)

			response := w.Result()

			if response.StatusCode != tt.outStatusCode {
				t.Errorf("Received Status code %d does not match expected status code %d", response.StatusCode, tt.outStatusCode)
			}
			var respBody interface{}
			// 에러가 발생한 경우 응답을 errMsg 타입으로 변환
			if tt.inErr != nil {
				var errmsg errMSG
				json.NewDecoder(response.Body).Decode(&errmsg)
				// 에러 메시지를 respBody 에 저장
				respBody = errmsg
			} else {
				// 에러가 없을 경우 응답을 product 타입의 슬라이스로 변환
				products := []models.Product{}
				json.NewDecoder(response.Body).Decode(&products)
				// 디코딩한 상품목록을 respBody 에 저장
				respBody = products
			}
			if !reflect.DeepEqual(respBody, tt.expectedRespBody) {
				t.Errorf("Received HTTP response body %+v does not match expected HTTP response Body %+v", respBody, tt.expectedRespBody)
			}
		})

	}
}
