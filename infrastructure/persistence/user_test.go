package persistence

import (
	"database/sql"
	"testing"

	"github.com/135yshr/go-confernce-mini-2022/domain/model"

	_ "github.com/go-sql-driver/mysql"
)

func NewTestDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:secret@tcp(localhost:3307)/smacare_ut?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestUser(t *testing.T) {
	t.Run("ユーザーを新しく登録したときデータベースに保存できること", func(t *testing.T) {
		db, err := NewTestDB()
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		if _, err := db.Exec("DELETE FROM users"); err != nil {
			t.Fatal(err)
		}

		user := &model.User{
			Name:     "山田太郎",
			Furigana: "やまだたろう",
		}

		sut := NewUserRepository(db)
		if err := sut.Insert(user); err != nil {
			t.Fatal(err)
		}

		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			t.Fatal(err)
		}
		defer rows.Close()

		users := make([]*model.User, 0)
		for rows.Next() {
			u := &model.User{}
			if err := rows.Scan(&u.ID, &u.Name, &u.Furigana, &u.Photo, &u.CreatedAt, &u.UpdatedAt); err != nil {
				t.Fatal(err)
			}
			users = append(users, u)
		}

		if rows.Err() != nil {
			t.Fatal(rows.Err())
		}

		if len(users) != 1 {
			t.Fatalf("ユーザーが1件ではありません: %d", len(users))
		}

		if users[0].Name != "山田太郎" {
			t.Fatalf("ユーザーの名前が正しくありません: %s", users[0].Name)
		}

		if users[0].Furigana != "やまだたろう" {
			t.Fatalf("ユーザーのふりがなが正しくありません: %s", users[0].Name)
		}
	})
}
