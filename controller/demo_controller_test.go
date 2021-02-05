package controller_test

import (
	"github.com/gin-gonic/gin"
	"github.com/steinfletcher/apitest"
	jsonPath "github.com/steinfletcher/apitest-jsonpath"
	"go-mock-best-practice/controller"
	"net/http"
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
		// when then
		//---------------------------------------
		apitest.New().
			Handler(router).
			Getf("/message?keyA=%s&url_long=1%s", "valA", "123456").
			Header("Client-Type", "pc").
			Cookie("sid", "id001").
			JSON(nil).
			Expect(t).
			Status(http.StatusOK).
			Assert(jsonPath.Equal(`$.code`, float64(2000))).
			Assert(jsonPath.Equal(`$.msg`, "Hello Mike!")).
			Body(`{"code":2000,"msg":"Hello Mike!"}`).
			End()
	})
}
