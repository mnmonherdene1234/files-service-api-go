package api

import (
	"encoding/json"
	"net/http"
)

type Map map[string]any

func JSONResponse(w http.ResponseWriter, statusCode int, data Map) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	bytes, err := json.Marshal(data)
	if err != nil {
		return
	}

	if _, err := w.Write(bytes); err != nil {
		return
	}
}
