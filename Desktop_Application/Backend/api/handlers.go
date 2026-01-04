package Api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type SearchedUser struct {
	Email      string
	Username   string
	Created_at string
	Online     bool
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("query")
	fmt.Println("query:", query)

	if len(query) < 3 {
		http.Error(w, "query too short or doesn't exist", http.StatusBadRequest)
		return
	}

	endpoint := "http://localhost:8090/search?q=" + url.QueryEscape(query)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+UserJWT)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "backend error", resp.StatusCode)
		return
	}

	var results []SearchedUser
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)

}
