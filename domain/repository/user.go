package repository

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE

import (
    "github.com/135yshr/go-confernce-mini-2022/domain/model"
)

type UserRepository interface {
    Insert(*model.User) error
}

