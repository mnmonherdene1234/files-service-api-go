package api

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	path := getURLPath(r.URL.Path)

	switch path {
	case "/files/upload":
		break
	default:
		JSONResponse(w, http.StatusNotFound, []byte(`{"message": "Not found"}`))
		break
	}
}

func getURLPath(path string) string {
	length := len(path)

	if length == 0 {
		return ""
	}

	if path[length-1] == '/' {
		return path[0 : length-1]
	}

	return path
}
