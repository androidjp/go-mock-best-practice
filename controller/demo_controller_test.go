package controller_test

import (
	"bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"github.com/steinfletcher/apitest"
	jsonPath "github.com/steinfletcher/apitest-jsonpath"
	"go-mock-best-practice/controller"
	"net/http"
	"reflect"
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

	// 通过mock API接口，来达到“想响应啥，就响应啥”的目的
	t.Run("should return `Hello Jasper` when mock the API", func(t *testing.T) {

		//---------------------------------------
		// given
		//---------------------------------------
		gin.SetMode(gin.TestMode)
		router := gin.New()
		//被mock掉了的接口
		router.GET("/message", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"code": 2000,
				"msg":  "Hello Jasper",
			})
		})

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
			Assert(jsonPath.Equal(`$.msg`, "Hello Jasper")).
			Body(`{"code":2000,"msg":"Hello Jasper"}`).
			End()
	})

	// 或者通过monkey.PatchInstanceMethod 来mock controller层逻辑，来达到目录
	t.Run("should return `Hello Jasper` when mock the controller layer logic", func(t *testing.T) {

		//---------------------------------------
		// given
		//---------------------------------------
		gin.SetMode(gin.TestMode)
		router := gin.New()
		demoCtrl := &controller.DemoController{}
		monkey.PatchInstanceMethod(reflect.TypeOf(demoCtrl), "GetMessage", func(ctrl *controller.DemoController, ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"code": 2000,
				"msg":  "Hello Jasper",
			})
		})
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
			Assert(jsonPath.Equal(`$.msg`, "Hello Jasper")).
			Body(`{"code":2000,"msg":"Hello Jasper"}`).
			End()
	})
}
