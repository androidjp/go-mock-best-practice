package controller_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-mock-best-practice/controller"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDemoController_GetMessage(t *testing.T) {
	t.Run("should return `Hello Mike` when do not mock the API", func(t *testing.T) {

		//---------------------------------------
		// given
		//---------------------------------------
		gin.SetMode(gin.TestMode)
		router := gin.New()
		demoCtrl := &controller.DemoController{}
		//待测试的接口
		router.GET("/message", demoCtrl.GetMessage)

		//---------------------------------------
		// when
		//---------------------------------------
		// 构建返回值
		w := httptest.NewRecorder()
		// 构建请求
		req, _ := http.NewRequest("GET", "/message?keyA=valA&url_long=123456", nil)
		//调用请求接口
		router.ServeHTTP(w, req)

		resp := w.Result()
		body, err := ioutil.ReadAll(resp.Body)

		//---------------------------------------
		// then
		//---------------------------------------
		assert.Nil(t, err)
		assert.Equal(t, "Hello Mike!", string(body))
	})
}
