package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port = "8081"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Привет Лилия"})
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Приложение развернуто, готово принимать запросы по url api/hello")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on :8080")
	http.HandleFunc("/api/hello", helloHandler)

	apiURL := fmt.Sprintf("http://localhost:%s/api/hello", port)
	fmt.Printf("Сервер запущен, url api = %s, готов принимать апи запросы\n", apiURL)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
