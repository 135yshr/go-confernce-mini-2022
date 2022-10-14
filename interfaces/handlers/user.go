package handlers

import (
	"net/http"

    "github.com/135yshr/go-confernce-mini-2022/domain/model"
	"github.com/135yshr/go-confernce-mini-2022/usecase"
)

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	uc usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) UserHandler {
	return &userHandler{uc: uc}
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
    h.uc.Create(&model.User{})
    w.WriteHeader(http.StatusCreated)
}
