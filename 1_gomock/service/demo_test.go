package gomock_service_test

import (
	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/assert"
	gomock_service "go-mock-best-practice/1_gomock/service"
	"testing"
)

func TestDemoService_CheckConnect(t *testing.T) {
	t.Run("全局变量打桩", func(t *testing.T) {
		// given
		stubs := gostub.Stub(&gomock_service.GlobalCount, 10).
			Stub(&gomock_service.Host, "https://www.bing.cn")
		defer stubs.Reset()
		demoSvr := &gomock_service.DemoService{Count: 8}

		// when then
		assert.Equal(t, "10:https://www.bing.cn:8", demoSvr.CheckConnect())
	})
}

func TestExec(t *testing.T) {
	t.Run("函数打桩", func(t *testing.T) {
		// given
		// 方式一：StubFunc
		stubs := gostub.StubFunc(&gomock_service.Exec, "xxx-vethName100-yyy", nil)
		defer stubs.Reset()
		// 方式二：Stub
		//stubs := gostub.Stub(&gomock_service.Exec, func(cmd string, args ...string) (string, error) {
		//	return "xxx-vethName100-yyy", nil
		//})
		//defer stubs.Reset()

		// when
		result, err := gomock_service.Exec("hello")

		// then
		assert.Equal(t, "xxx-vethName100-yyy", result)
		assert.Nil(t, err)
	})
}

func TestDestroyResource(t *testing.T) {
	t.Run("过程打桩", func(t *testing.T) {
		// given
		// 方式一：StubFunc（当你想这个过程啥都不做时，可以这样）
		stubs := gostub.StubFunc(&gomock_service.DestroyResource)
		defer stubs.Reset()
		// 方式二：Stub
		//stubs := gostub.Stub(&gomock_service.DestroyResource, func() {
		//	// do sth
		//})
		//defer stubs.Reset()
	})
}
