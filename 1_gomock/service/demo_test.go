package gomock_service_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	gomock_service "go-mock-best-practice/1_gomock/service"
	mock_gomock_db "go-mock-best-practice/test/1_gomock/db"
	"testing"
)

func TestDemoService_InsertData(t *testing.T) {
	Convey("should return tips and err", t, func() {
		Convey("when repo return `db connection error`", func() {
			// given
			// 1. 初始化 mock控制器
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// 2. 初始化mock对象，并注入控制器
			mockRepo := mock_gomock_db.NewMockRepository(ctrl)

			// 3. 设定mock对象的返回值
			mockRepo.EXPECT().Create("name", []byte("jasper")).Return(errors.New("db connection error")).Times(3)
			demoSvr := &gomock_service.DemoService{Repo: mockRepo}

			// when
			data, err := demoSvr.InsertData("name", "jasper")

			// then
			So(data, ShouldEqual, "need to retry")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "db connection error")
		})
	})

	Convey("should return empty and err", t, func() {
		Convey("when repo return other error", func() {
			// given
			// 1. 初始化 mock控制器
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// 2. 初始化mock对象，并注入控制器
			mockRepo := mock_gomock_db.NewMockRepository(ctrl)

			// 3. 设定mock对象的返回值
			mockRepo.EXPECT().Create("name", []byte("jasper")).Return(errors.New("数据不存在")).Times(3)
			demoSvr := &gomock_service.DemoService{Repo: mockRepo}

			// when
			data, err := demoSvr.InsertData("name", "jasper")

			// then
			So(data, ShouldEqual, "")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "数据不存在")
		})
	})
	Convey("should return success and nil", t, func() {
		Convey("when repo return nil", func() {
			// given
			// 1. 初始化 mock控制器
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// 2. 初始化mock对象，并注入控制器
			mockRepo := mock_gomock_db.NewMockRepository(ctrl)

			// 3. 设定mock对象的返回值
			mockRepo.EXPECT().Create("name", []byte("jasper")).Return(nil)
			demoSvr := &gomock_service.DemoService{Repo: mockRepo}

			// when
			data, err := demoSvr.InsertData("name", "jasper")

			// then
			assert.Equal(t, "success", data)
			assert.Nil(t, err)
			So(data, ShouldEqual, "success")
			So(err, ShouldBeNil)
		})
		Convey("when repo return nil at the third time", func() {
			// given
			// 1. 初始化 mock控制器
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// 2. 初始化mock对象，并注入控制器
			mockRepo := mock_gomock_db.NewMockRepository(ctrl)

			// 3. 设定mock对象的返回值
			mockRepo.EXPECT().Create("name", []byte("jasper")).Return(errors.New("db connection error")).Times(2) // 前两次返回错误
			mockRepo.EXPECT().Create("name", []byte("jasper")).Return(nil)                                        // 第三次正常

			demoSvr := &gomock_service.DemoService{Repo: mockRepo}

			// when
			data, err := demoSvr.InsertData("name", "jasper")

			// then
			So(data, ShouldEqual, "success")
			So(err, ShouldBeNil)
		})
	})
}

func TestDemoService_CheckAndUpdateData(t *testing.T) {

	Convey("should return nil", t, func() {
		Convey("when can find existed data", func() {
			// given
			// 1. 初始化 mock控制器
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// 2. 初始化mock对象，并注入控制器
			mockRepo := mock_gomock_db.NewMockRepository(ctrl)

			// 3. 设定mock对象的返回值
			retrieveName := mockRepo.EXPECT().Retrieve("name").Return([]byte("jasper"), nil)
			mockRepo.EXPECT().Update("name", []byte("mike")).Return(nil).After(retrieveName) // update 在 retrieve 之后

			// 指定调用顺序的方式二
			//gomock.InOrder(
			//	mockRepo.EXPECT().Retrieve("name").Return([]byte("jasper"), nil),
			//	mockRepo.EXPECT().Update("name", []byte("mike")).Return(nil),
			//)

			demoSvr := &gomock_service.DemoService{Repo: mockRepo}

			// when
			err := demoSvr.CheckAndUpdateData("name", "mike")

			// then
			So(err, ShouldBeNil)
		})
	})
}

func TestDemoService_CheckConnect(t *testing.T) {
	Convey("全局变量打桩", t, func() {
		// given
		stubs := gostub.Stub(&gomock_service.GlobalCount, 10).
			Stub(&gomock_service.Host, "https://www.bing.cn")
		defer stubs.Reset()
		demoSvr := &gomock_service.DemoService{Count: 8}

		// when then
		So(demoSvr.CheckConnect(), ShouldEqual, "10:https://www.bing.cn:8")
	})
}

func TestExec(t *testing.T) {
	Convey("函数打桩", t, func() {
		// given
		stubs := gostub.Stub(&gomock_service.Exec, func(cmd string, args ...string) (string, error) {
			return "xxx-vethName100-yyy", nil
		})
		defer stubs.Reset()

		// when
		result, err := gomock_service.Exec("hello")

		// then
		So(result, ShouldEqual, "xxx-vethName100-yyy")
		So(err, ShouldBeNil)
	})
}

func TestDestroyResource(t *testing.T) {
	Convey("过程打桩", t, func() {
		// given
		stubs := gostub.StubFunc(&gomock_service.DestroyResource)
		defer stubs.Reset()
	})
}
