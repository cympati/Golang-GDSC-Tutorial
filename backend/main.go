package main

import (
	"encoding/json"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		responseValue := map[string]any{"Message": "Hi, GDSC3 :)"}
		response, _ := json.Marshal(responseValue)
		w.Write(response)
	})

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
