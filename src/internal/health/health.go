package health

import (
	"encoding/json"
	"net/http"
	"time"
)

func Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"time":   time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
	})
}
