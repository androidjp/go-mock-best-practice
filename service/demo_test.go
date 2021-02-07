package service

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	"go-mock-best-practice/adapter"
	"regexp"
	"testing"
)

// 直接对入参db对象进行mock
func TestDemoService_AddStudentDirectly(t *testing.T) {
	Convey("should return Student{id:1, name:`mike`}", t, func() {
		Convey("when input name `mike` and can insert mysql db successfully", func() {
			//---------------------------------------
			// given
			//---------------------------------------
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(`insert into students(name) values(?)`)).WithArgs("mike").WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectExec(regexp.QuoteMeta("insert into classroom_1(stu_id) values(?)")).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			//---------------------------------------
			// when
			//---------------------------------------
			svr := &DemoService{}
			stu, err := svr.AddStudentDirectly(db, "mike")

			//---------------------------------------
			// then
			//---------------------------------------
			So(err, ShouldBeNil)
			So(stu, ShouldNotBeNil)
			So(stu.ID, ShouldEqual, 1)
			So(stu.Name, ShouldEqual, "mike")
		})
	})
}

// 对潜在的db对象进行mock，间接地mock了repository层对象（当然，这种方式其实是不这么优雅地，理论，使用gomock等直接mock掉repository层对象，才是更佳的实践，这里只是展示如何使用sqlmock）
func TestDemoService_AddStudentByName(t *testing.T) {
	Convey("should return nil", t, func() {
		Convey("when input name `mike` and can insert mysql db successfully", func() {
			//---------------------------------------
			// given
			//---------------------------------------
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			// 注意：这里是打桩的关键：将mock的db对象，作为Open函数的返回
			stubs := gostub.StubFunc(&adapter.Open, db, nil)
			defer stubs.Reset()

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(`insert into students(name) values(?)`)).WithArgs("mike").WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectExec(regexp.QuoteMeta("insert into classroom_1(stu_id) values(?)")).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			//---------------------------------------
			// when
			//---------------------------------------
			svr := &DemoService{}
			err = svr.AddStudentByName("mike")

			//---------------------------------------
			// then
			//---------------------------------------
			So(err, ShouldBeNil)
		})
	})
}
