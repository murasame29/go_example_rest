package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go_example_crud/models"
	"net/http"
	"strconv"
)

type DB struct {
	DB *sql.DB
}

func NewDB(db *sql.DB) *DB {
	return &DB{
		DB: db,
	}
}

// 全てを取得
func (db *DB) GetAll(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query(`SELECT id, name ,age FROM users`)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	users := []models.Users{}

	for rows.Next() {
		user := models.Users{}

		if err := rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		users = append(users, user)
	}

	body, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

// ユーザの挿入
func (db *DB) Insert(w http.ResponseWriter, r *http.Request) {
	user := models.Users{}
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	age, err := strconv.Atoi(r.Form.Get("age"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//パスワードのハッシュ化は無し
	insert, err := db.DB.Prepare("INSERT INTO users (name,age,password) VALUES ($1,$2,$3) RETURNING id")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := insert.QueryRow(r.Form.Get("name"), age, r.Form.Get("passowrd")).Scan(&user.Id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"message":"success","id":%d}`, user.Id)))
}

// ユーザの更新
func (db *DB) Update(w http.ResponseWriter, r *http.Request) {

}

// 全ユーザの削除
func (db *DB) DeleteAll(w http.ResponseWriter, r *http.Request) {
	if _, err := db.DB.Exec(`TRUNCATE TABLE users`); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"message":"success"}`)))
}
