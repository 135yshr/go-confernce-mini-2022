package usecase

import (
    "testing"

    "github.com/golang/mock/gomock"
    "github.com/135yshr/go-confernce-mini-2022/domain/model"
    repository "github.com/135yshr/go-confernce-mini-2022/domain/repository/mock"
)

func TestUser(t *testing.T) {
    t.Run("利用者ユースケース", func(t *testing.T) {
        t.Run("ユースケースが正常に終了すること", func(t *testing.T) {
            ctrl := gomock.NewController(t)
            defer ctrl.Finish()

            repo := repository.NewMockUserRepository(ctrl)
            repo.EXPECT().Insert(gomock.Any()).Return(nil)

            sut := NewUserUsecase(repo)

            if err := sut.Create(&model.User{}); err != nil {
                t.Errorf("error should be nil, but got %v", err)
            }
        })
    })
}
