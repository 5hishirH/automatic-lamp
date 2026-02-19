package api

import (
	"encoding/json"
	"net/http"
)

// Response defines the structure of our JSON output
type Response struct {
	Message string `json:"message"`
	Runtime string `json:"runtime"`
	Status  int    `json:"status"`
}

// Handler is the entry point for Vercel Serverless Functions.
// It MUST be named 'Handler' and have this exact signature.
func Handler(w http.ResponseWriter, r *http.Request) {
	// 1. Set the response header to JSON
	w.Header().Set("Content-Type", "application/json")

	// 2. Simple routing: restrict this endpoint to GET requests only
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Method not allowed",
		})
		return
	}

	// 3. Construct the success response
	resp := Response{
		Message: "Hello from a lightning-fast Go API!",
		Runtime: "Golang on Vercel",
		Status:  http.StatusOK,
	}

	// 4. Encode the struct to JSON and write it to the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
