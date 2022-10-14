package handlers

import (
	"log"
	"net/http"

	"github.com/135yshr/go-confernce-mini-2022/domain/model"
	"github.com/135yshr/go-confernce-mini-2022/usecase"
)

type UserHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	uc usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) UserHandler {
	return &userHandler{uc: uc}
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("handler: serve http")
	h.Create(w, r)
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("handler: create user")
	user := &model.User{
		Name:     "山田太郎",
		Furigana: "ヤマダタロウ",
	}
	h.uc.Create(user)
	w.WriteHeader(http.StatusCreated)
}
