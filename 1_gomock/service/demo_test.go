package gomock_service_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/assert"
	gomock_service "go-mock-best-practice/1_gomock/service"
	mock_gomock_db "go-mock-best-practice/test/1_gomock/db"
	"testing"
)

func TestDemoService_InsertData(t *testing.T) {
	t.Run("should_return_tips_and_err_given_InsertData_when_repo_return_mongo_connection_error", func(t *testing.T) {
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
		assert.Equal(t, "need to retry", data)
		assert.NotNil(t, err)
		assert.Equal(t, "db connection error", err.Error())
	})

	t.Run("should_return_empty_and_err_given_InsertData_when_repo_return_other_error", func(t *testing.T) {
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
		assert.Equal(t, "", data)
		assert.NotNil(t, err)
		assert.Equal(t, "数据不存在", err.Error())
	})

	t.Run("should_return_success_given_InsertData_when_repo_return_nil", func(t *testing.T) {
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
	})

	t.Run("should_return_success_given_InsertData_when_repo_return_nil_at_the_third_time", func(t *testing.T) {
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
		assert.Equal(t, "success", data)
		assert.Nil(t, err)
	})
}

func TestDemoService_CheckAndUpdateData(t *testing.T) {
	t.Run("should_return_nil_given_can_find_existed_data", func(t *testing.T) {
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
		assert.Nil(t, err)
	})
}

func TestDemoService_CheckConnect(t *testing.T) {
	t.Run("全局变量打桩", func(t *testing.T) {
		// given
		stubs := gostub.Stub(&gomock_service.GlobalCount, 10).
			Stub(&gomock_service.Host, "https://www.bing.cn")
		defer stubs.Reset()
		demoSvr := &gomock_service.DemoService{}

		// when then
		assert.Equal(t, "10:https://www.bing.cn:0", demoSvr.CheckConnect())
	})
}

func TestExec(t *testing.T) {
	t.Run("函数打桩", func(t *testing.T) {
		// given
		stubs := gostub.Stub(&gomock_service.Exec, func(cmd string, args ...string) (string, error) {
			return "xxx-vethName100-yyy", nil
		})
		defer stubs.Reset()

		// when
		result, err := gomock_service.Exec("hello")

		// then
		assert.Equal(t, "xxx-vethName100-yyy", result)
		assert.Nil(t, err)
	})
}
