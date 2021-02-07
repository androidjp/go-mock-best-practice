package repository_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	"go-mock-best-practice/adapter"
	"go-mock-best-practice/repository"
	"regexp"
	"testing"
)

func TestMySQLRepository_CreateStudent(t *testing.T) {
	Convey("should return error `db connect error`", t, func() {
		Convey("when occurs error during db exec", func() {
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
			sqlRepository := repository.NewMySQLRepository()
			student, err := sqlRepository.CreateStudent("mike")

			//---------------------------------------
			// then
			//---------------------------------------

			So(err, ShouldBeNil)
			So(student, ShouldNotBeNil)
			So(student.ID, ShouldEqual, 1)
			So(student.Name, ShouldEqual, "mike")
		})
	})
}
