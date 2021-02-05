package repository_test

import (
	. "bou.ke/monkey"
	. "github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"go-mock-best-practice/entities"
	"go-mock-best-practice/mock/mock_repository"
	"go-mock-best-practice/repository"
	"reflect"
	"testing"
)

func TestRepo_Retrieve(t *testing.T) {
	Convey("[桩中桩案例]", t, func() {
		Convey("当我需要mock一个接口，并且，重新定义这个mock对象的某个方法时，使用", func() {
			//---------------------------------------
			// given
			//---------------------------------------
			// 接口打桩
			ctrl := NewController(t)
			defer ctrl.Finish()
			mockRepo := mock_repository.NewMockRepository(ctrl)
			// 函数打桩
			Patch(repository.GetInstance, func() repository.Repository {
				return mockRepo
			})
			defer UnpatchAll()
			// 成员方法打桩
			PatchInstanceMethod(reflect.TypeOf(mockRepo), "Retrieve", func(repo *mock_repository.MockRepository, name string, movie *entities.Movie) error {
				*movie = entities.Movie{Name: name, Type: "Love", Score: 95}
				return nil
			})
			//---------------------------------------
			// when
			//---------------------------------------
			repo := repository.GetInstance()
			var movie = &entities.Movie{}
			err := repo.Retrieve("Titanic", movie)

			//---------------------------------------
			// then
			//---------------------------------------
			So(err, ShouldBeNil)
			So(movie.Name, ShouldEqual, "Titanic")
			So(movie.Type, ShouldEqual, "Love")
			So(movie.Score, ShouldEqual, 95)
		})
	})
}
