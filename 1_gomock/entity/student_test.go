package entity_test

import (
	"errors"
	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/assert"
	"go-mock-best-practice/1_gomock/adapter"
	"go-mock-best-practice/1_gomock/entity"
	"testing"
)

func TestStudent_Print(t *testing.T) {
	t.Run("第三方库函数打桩[成功情况]", func(t *testing.T) {
		// given
		var mikeStr = `{"name":"Jasper", "age":18}`
		stubs := gostub.StubFunc(&adapter.Marshal, []byte(mikeStr), nil)
		defer stubs.Reset()

		stu := &entity.Student{Name: "Mike", Age: 18}

		// when
		res, err := stu.Print()

		// then
		assert.Equal(t, "{\"name\":\"Jasper\", \"age\":18}", res)
		assert.Nil(t, err)
	})

	t.Run("第三方库函数打桩[失败情况]", func(t *testing.T) {
		// given
		stubs := gostub.StubFunc(&adapter.Marshal, nil, errors.New("error"))
		defer stubs.Reset()

		stu := &entity.Student{Name: "Mike", Age: 18}

		// when
		res, err := stu.Print()

		// then
		assert.Equal(t, "", res)
		assert.Equal(t, "error", err.Error())
	})
}
