package main

import (
	"cympati/Golang-GDSC-Tutorial/config"
	"cympati/Golang-GDSC-Tutorial/handler"
	"cympati/Golang-GDSC-Tutorial/repository"
	"cympati/Golang-GDSC-Tutorial/service"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

func main() {
	print("Hi, GDSC1 :)")

	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}

	db, err := sql.Open("mysql", config.C.DB_HOST) // Like this
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewRepositoryDB(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/signup", userHandler.SignUp)
	http.HandleFunc("/signin", userHandler.SignIn)
	http.HandleFunc("/listuser", userHandler.ListUsers)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		responseValue := map[string]any{"Message": "Hi, GDSC3 :)"}
		response, _ := json.Marshal(responseValue)
		w.Write(response)
	})

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}

	defer db.Close()

}
