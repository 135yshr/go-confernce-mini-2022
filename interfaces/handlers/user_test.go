package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

	usecase "github.com/135yshr/go-confernce-mini-2022/usecase/mock"
)

func TestUserHandler(t *testing.T) {
	t.Run("UserHandler", func(t *testing.T) {
		t.Run("利用者情報を受け取り正常に処理が終了すること", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uc := usecase.NewMockUserUsecase(ctrl)
			uc.EXPECT().Create(gomock.Any()).Return(nil)

			req, err := http.NewRequest("POST", "/api/v1/users", nil)
			if err != nil {
				t.Fatal(err)
			}

			sut := NewUserHandler(uc)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(sut.Create)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusCreated {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
		})
	})
}
