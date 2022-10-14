package usecase

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE

import (
	"github.com/135yshr/go-confernce-mini-2022/domain/model"
	"github.com/135yshr/go-confernce-mini-2022/domain/repository"
)

type UserUsecase interface {
	Create(*model.User) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (uc *userUsecase) Create(u *model.User) error {
	if err := uc.repo.Insert(u); err != nil {
		return NewError(100, "failed to insert user")
	}
	return nil
}
