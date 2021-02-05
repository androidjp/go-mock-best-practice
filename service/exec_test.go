package service_test

import (
	. "bou.ke/monkey"
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"go-mock-best-practice/service"
	"reflect"
	"testing"
)

func TestExec(t *testing.T) {
	// 直接就能进行打桩，不需要更改相关的全局函数为函数对象
	Convey("【函数打桩】should has dight", t, func() {
		Convey("when successful", func() {
			//---------------------------------------
			// given
			//---------------------------------------
			guard := Patch(service.Exec, func(cmd string, args ...string) (string, error) {
				return "sss", nil
			})
			defer guard.Unpatch()

			//---------------------------------------
			// when
			//---------------------------------------
			output, err := service.Exec("cmd01", "--conf", "conf/app_test.yaml")

			//---------------------------------------
			// then
			//---------------------------------------
			So(output, ShouldEqual, "sss")
			So(err, ShouldBeNil)
		})
	})
}

func TestInternalDoSth(t *testing.T) {
	Convey("【过程打桩】should return valB", t, func() {
		Convey("when get keyA", func() {
			//---------------------------------------
			// given
			//---------------------------------------

			patchGuard := Patch(service.InternalDoSth, func(mData map[string]interface{}) {
				mData["keyA"] = "valB"
			})
			defer patchGuard.Unpatch()

			//---------------------------------------
			// when
			//---------------------------------------
			mData := map[string]interface{}{}
			service.InternalDoSth(mData)

			//---------------------------------------
			// then
			//---------------------------------------
			So(mData["keyA"], ShouldEqual, "valB")
		})
	})
}

func TestEtcd_Get(t *testing.T) {
	Convey("should [方法打桩]", t, func() {
		Convey("when fdfdf", func() {
			//---------------------------------------
			// given
			//---------------------------------------
			var e = &service.Etcd{}
			guard := PatchInstanceMethod(reflect.TypeOf(e), "Get", func(e *service.Etcd, id int) []string {
				return []string{"Jasper"}
			})
			defer guard.Unpatch()

			//---------------------------------------
			// when
			//---------------------------------------
			res := e.Get(1)

			//---------------------------------------
			// then
			//---------------------------------------
			So(res[0], ShouldEqual, "Jasper")
		})
	})
}

func TestEtcd_GetAndSave(t *testing.T) {
	Convey("should [同时打桩多个方法]", t, func() {
		Convey("when GetAndSave 方法 先后调用 Get 方法 和 Save 方法", func() {
			//---------------------------------------
			// given
			//---------------------------------------
			var e = &service.Etcd{}
			// stub Get
			theVals := make([]string, 0)
			theVals = append(theVals, "A")
			PatchInstanceMethod(reflect.TypeOf(e), "Get", func(e *service.Etcd, id int) []string {
				return theVals
			})
			// stub Save
			PatchInstanceMethod(reflect.TypeOf(e), "Save", func(e *service.Etcd, vals []string) (string, error) {
				return "", errors.New("occurs error")
			})

			// 一键删除所有补丁
			defer UnpatchAll()

			//---------------------------------------
			// when
			//---------------------------------------
			res, err := e.GetAndSave(0)

			//---------------------------------------
			// then
			//---------------------------------------
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "occurs error")
			So(res, ShouldBeEmpty)
			So(theVals[0], ShouldEqual, "C")
		})
	})
	Convey("should [同时打桩多个方法，连自己都打]", t, func() {
		Convey("when GetAndSave 方法 先后调用 Get 方法 和 Save 方法", func() {
			//---------------------------------------
			// given
			//---------------------------------------
			var e = &service.Etcd{}

			// stub GetAndSave
			PatchInstanceMethod(reflect.TypeOf(e), "GetAndSave", func(e *service.Etcd, id int) (string, error) {
				vals := e.Get(id)
				if vals[0] == "A" {
					vals[0] = "D"
				}
				return e.Save(vals)
			})
			// stub Get
			theVals := make([]string, 0)
			theVals = append(theVals, "A")
			PatchInstanceMethod(reflect.TypeOf(e), "Get", func(e *service.Etcd, id int) []string {
				return theVals
			})
			// stub Save
			PatchInstanceMethod(reflect.TypeOf(e), "Save", func(e *service.Etcd, vals []string) (string, error) {
				return "", errors.New("occurs error")
			})

			// 一键删除所有补丁
			defer UnpatchAll()

			//---------------------------------------
			// when
			//---------------------------------------
			res, err := e.GetAndSave(0)

			//---------------------------------------
			// then
			//---------------------------------------
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "occurs error")
			So(res, ShouldBeEmpty)
			So(theVals[0], ShouldEqual, "D")
		})
	})

}
