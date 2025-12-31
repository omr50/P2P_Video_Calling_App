package api

import (
	"encoding/json"
	"net/http"
)

func EmailSearch(w http.ResponseWriter, r *http.Request) {
	partialEmail := r.URL.Query().Get("q")

	if partialEmail == "" {
		http.Error(w, "missing search query", http.StatusBadRequest)
	}

	// call db query, should return back some sort of array with the results?
	results, err := FindByPartialEmail(partialEmail)

	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	ctx := r.Context()
	for i := range results {
		redisEmailKey := "online:user:" + results[i].Email

		// search redis for this
		exists, err := RedisClient.Exists(ctx, redisEmailKey).Result()
		if err != nil {
			continue
		}
		results[i].Online = (exists == 1)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
