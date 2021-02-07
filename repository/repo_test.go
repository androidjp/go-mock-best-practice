package repository_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	"go-mock-best-practice/adapter"
	"go-mock-best-practice/repository"
	"regexp"
	"testing"
	"xorm.io/xorm"
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
			mockEngine, err := xorm.NewEngine("mysql", "root:123@/test?charset=utf8mb4")
			mockEngine.DB().DB = db

			// 注意：这里是打桩的关键：将mock的db对象，作为Open函数的返回
			stubs := gostub.StubFunc(&adapter.Open, mockEngine, nil)
			defer stubs.Reset()

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `students` (`name`) VALUES (?)")).WithArgs("mike").WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `classroom_1` (`stu_id`) VALUES (?)")).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
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

func TestMySQLRepository_GetStudents(t *testing.T) {
	Convey("should return `Jim` and `Jimmy`", t, func() {
		Convey("when find by key `Ji%` and limit 2", func() {
			//---------------------------------------
			// given
			//---------------------------------------
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			mockEngine, err := xorm.NewEngine("mysql", "root:123@/test?charset=utf8mb4")
			mockEngine.DB().DB = db

			// 注意：这里是打桩的关键：将mock的db对象，作为Open函数的返回
			stubs := gostub.StubFunc(&adapter.Open, mockEngine, nil)
			defer stubs.Reset()

			rows := sqlmock.NewRows([]string{"name"}).
				AddRow("Jim").
				AddRow("Jimmy")

			mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM `students` WHERE (name LIKE ?) LIMIT 2")).WillReturnRows(rows)

			//---------------------------------------
			// when
			//---------------------------------------
			students, err := repository.NewMySQLRepository().GetStudents("Ji%", 2)

			//---------------------------------------
			// then
			//---------------------------------------
			So(err, ShouldBeNil)
			So(students, ShouldNotBeNil)
			So(students, ShouldHaveLength, 2)
			So(students[0].Name, ShouldEqual, "Jim")
			So(students[1].Name, ShouldEqual, "Jimmy")
			So(students[0].ID, ShouldEqual, 0)
			So(students[1].ID, ShouldEqual, 0)
		})
	})
}

//func TestNewMySQLRepository(t *testing.T) {
//	student, err := repository.NewMySQLRepository().CreateStudent("jasper")
//	fmt.Printf("%v\n", student)
//	fmt.Printf("%v\n", err)
//}
