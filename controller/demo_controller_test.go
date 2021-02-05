package controller_test

import (
	"bou.ke/monkey"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-mock-best-practice/controller"
	"gopkg.in/h2non/gock.v1"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestDemoController_GetMessage(t *testing.T) {
	t.Run("should return `Hello Mike` when do not mock the API", func(t *testing.T) {

		//---------------------------------------
		// given
		//---------------------------------------
		demoCtrl := &controller.DemoController{}
		ts := httptest.NewServer(http.HandlerFunc(demoCtrl.GetMessage))
		defer ts.Close()

		//---------------------------------------
		// when
		//---------------------------------------
		resp, err := http.Get(ts.URL)
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		//---------------------------------------
		// then
		//---------------------------------------
		assert.Nil(t, err)
		assert.Equal(t, "Hello Mike!", string(bodyBytes))
	})

	// 通过mock API接口，来达到“想响应啥，就响应啥”的目的
	t.Run("should return `Hello Jasper` when mock the API", func(t *testing.T) {

		//---------------------------------------
		// given
		//---------------------------------------
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
			fmt.Fprintf(w, "Hello Jasper")
		}))
		defer ts.Close()

		//---------------------------------------
		// when
		//---------------------------------------
		resp, err := http.Get(ts.URL)
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		//---------------------------------------
		// then
		//---------------------------------------
		assert.Nil(t, err)
		assert.Equal(t, "Hello Jasper", string(bodyBytes))
	})

	// 或者通过monkey.PatchInstanceMethod 来mock controller层逻辑，来达到mock目的
	t.Run("should return `Hello Jasper` when mock the controller layer logic", func(t *testing.T) {

		//---------------------------------------
		// given
		//---------------------------------------
		demoCtrl := &controller.DemoController{}
		monkey.PatchInstanceMethod(reflect.TypeOf(demoCtrl), "GetMessage", func(ctrl *controller.DemoController, w http.ResponseWriter, request *http.Request) {
			fmt.Fprintf(w, "Hello Jasper")
		})
		ts := httptest.NewServer(http.HandlerFunc(demoCtrl.GetMessage))
		defer ts.Close()

		//---------------------------------------
		// when
		//---------------------------------------
		resp, err := http.Get(ts.URL)
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		//---------------------------------------
		// then
		//---------------------------------------
		assert.Nil(t, err)
		assert.Equal(t, "Hello Jasper", string(bodyBytes))
	})

	// 或者通过gock库，来达到mock目的
	// gock github地址：https://github.com/h2non/gock
	t.Run("should return `Hello Jasper` when mock api by gock", func(t *testing.T) {

		//---------------------------------------
		// given
		//---------------------------------------
		defer gock.Off()

		gock.New("http://localhost").
			Get("/message").
			Reply(200).
			JSON(map[string]interface{}{"code": 2000, "msg": "Hello Jasper"})

		//---------------------------------------
		// when
		//---------------------------------------
		resp, err := http.Get("http://localhost/message?keyA=valA&url_long=123456")
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		//---------------------------------------
		// then
		//---------------------------------------
		assert.Nil(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, "{\"code\":2000,\"msg\":\"Hello Jasper\"}\n", string(bodyBytes))
	})
}
