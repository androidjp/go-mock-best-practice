package entity_test

import (
	"errors"
	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	"go-mock-best-practice/1_gomock/adapter"
	"go-mock-best-practice/1_gomock/entity"
	"testing"
)

func TestStudent_Print(t *testing.T) {
	Convey("TestStudent_Print", t, func() {
		Convey("第三方库函数打桩[成功情况]", func() {
			// given
			var mikeStr = `{"name":"Jasper", "age":18}`
			stubs := gostub.StubFunc(&adapter.Marshal, []byte(mikeStr), nil)
			defer stubs.Reset()

			stu := &entity.Student{Name: "Mike", Age: 18}

			// when
			res, err := stu.Print()

			// then
			So(res, ShouldEqual, "{\"name\":\"Jasper\", \"age\":18}")
			So(err, ShouldBeNil)
		})

		Convey("第三方库函数打桩[失败情况]", func() {
			// given
			stubs := gostub.StubFunc(&adapter.Marshal, nil, errors.New("error"))
			defer stubs.Reset()

			stu := &entity.Student{Name: "Mike", Age: 18}

			// when
			res, err := stu.Print()

			// then
			So(res, ShouldEqual, "")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "error")
		})
	})
}
