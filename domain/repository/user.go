package repository

import (
    "github.com/135yshr/go-confernce-mini-2022/domain/model"
)

type UserRepository interface {
    Insert(*model.User) error
}

