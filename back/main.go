package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/save-user", saveUserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func saveUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")

	fmt.Println(r.Header)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Метод %s не поддерживается", r.Method)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Ошибка декодирования данных: %v", err)
		return
	}
	defer r.Body.Close()

	db, err := sql.Open("postgres", "user=user password=password dbname=postgres port=5432 sslmode=disable")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Не удалось подключиться к базе данных: %v", err)
		return
	}
	defer db.Close()

	_, err = db.Exec(
		"INSERT INTO users (username, password) VALUES ($1, $2)",
		user.Username,
		user.Password,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Ошибка вставки данных: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Пользователь успешно сохранён")
}
