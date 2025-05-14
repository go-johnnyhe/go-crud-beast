package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"math/rand"
	"strings"
)

type User struct {
	ID	string	`json:"id"`
	Name string	`json:"name"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})




	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"error": "Only POST method allowed"})
			return
		}	

		var newUser User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON"})
			return
		}

		if !strings.Contains(newUser.Email, "@") || newUser.Email == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid email format"})
			return
		}
		// validate here
		
		newUser.ID = fmt.Sprintf("usr_%d", rand.Intn(10000))
		
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)

	})

	log.Printf("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
