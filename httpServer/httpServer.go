package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

type Data struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	//curl -X POST localhost:8080/data -d '{ "name" : "hello" }'
	var data Data
	if r.Method == http.MethodPost {
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "проблема json", http.StatusBadRequest)
		}
		fmt.Printf("Получены данные: %v\n", data)
		fmt.Fprintln(w, "Данные приняты")
	} else {
		http.Error(w, "Поддержка только Post запроса", http.StatusMethodNotAllowed)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/data", dataHandler)
	loggedMux := logger(mux)
	// http.HandleFunc("/hello", helloHandler)
	// http.HandleFunc("/data", dataHandler)
	fmt.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", loggedMux))

}
