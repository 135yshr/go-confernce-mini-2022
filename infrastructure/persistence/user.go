package persistence

import (
	"database/sql"

	"github.com/135yshr/go-confernce-mini-2022/domain/model"
	"github.com/135yshr/go-confernce-mini-2022/domain/repository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db}
}

func (u *userRepository) Insert(user *model.User) error {
	stmt, err := u.db.Prepare("INSERT INTO users (name, furigana, photo) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Furigana, user.Photo)
	return err
}
