package gomock_service_test

import (
	"errors"
	"github.com/golang/mock/gomock"
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
		mockRepo.EXPECT().Create("name", []byte("jasper")).Return(errors.New("db connection error"))
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
		mockRepo.EXPECT().Create("name", []byte("jasper")).Return(errors.New("数据不存在"))
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

}
