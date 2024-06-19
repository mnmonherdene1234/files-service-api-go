package api

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	path := getURLPath(r.URL.Path)

	switch path {
	case "/upload":
		uploadHandler(w, r)
		break
	default:
		JSONResponse(w, http.StatusNotFound, Map{
			"message": "Not found",
		})
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

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, formFileErr := r.FormFile("file")

	if formFileErr != nil {
		JSONResponse(w, http.StatusInternalServerError, Map{
			"message": formFileErr,
		})
		return
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	f, openFileErr := os.OpenFile("/pathToStoreFile/"+fileHeader.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if openFileErr != nil {
		JSONResponse(w, http.StatusInternalServerError, Map{
			"message": openFileErr,
		})
		return
	}

	if _, err := io.Copy(f, file); err != nil {
		JSONResponse(w, http.StatusInternalServerError, Map{
			"message": err,
		})
		return
	}

	JSONResponse(w, http.StatusOK, Map{
		"message": "File uploaded",
	})
}
