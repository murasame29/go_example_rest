package handle

import (
	"go_example_crud/config"
	"go_example_crud/repository"
	"log"
	"net/http"
)

// 全てのユーザに対する処理
func UsersHandle(w http.ResponseWriter, r *http.Request) {
	db, err := config.Open()
	if err != nil {
		log.Fatal(err)
	}

	usersHandler := repository.NewDB(db)

	switch r.Method {

	case http.MethodGet: //一括取得

		usersHandler.GetAll(w, r)

	case http.MethodPost: //新規作成

		usersHandler.Insert(w, r)

	case http.MethodPut: //使わない

		usersHandler.Update(w, r)

	case http.MethodDelete: //全削除

		usersHandler.DeleteAll(w, r)

	default: //それ以外も使わない

		w.WriteHeader(http.StatusBadRequest)
	}
}
