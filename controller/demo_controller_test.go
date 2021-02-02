package controller_test

import (
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
}
