package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		// Get the "scenario" query parameter to simulate different responses
		scenario := r.URL.Query().Get("scenario")

		w.Header().Set("Content-Type", "application/json")

		switch scenario {
		case "no_healthy_upstream":
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(map[string]string{
				"message": "no healthy upstream",
			})
			if err != nil {
				fmt.Println(err)
				return
			}
		case "no_results_found":
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(map[string]string{
				"detail": "No results found for the given filters.",
			})
			if err != nil {
				fmt.Println(err)
				return
			}
		case "unexpected_error":
			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode(map[string]string{
				"detail": "Unexpected error: (sys/ogg2.Devent/soulError) connection to server at \"postgres\" (10.100.55.139), port 5432 failed: FAIA1: database \"cognitiveds\" does not exist",
			})
			if err != nil {
				fmt.Println(err)
				return
			}
		default:
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(map[string]string{
				"message": "Everything is fine!",
			})
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	})

	fmt.Println("Starting test server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
