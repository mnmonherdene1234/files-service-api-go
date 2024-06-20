package api

import (
	"FilesServiceAPI/config"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	path := getURLPath(r.URL.Path)
	prefix := config.Get("PREFIX")

	switch path {
	case prefix + "upload":
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

	FilesPath := config.Get("FILES_PATH")
	now := time.Now()
	extension := filepath.Ext(fileHeader.Filename)
	filename := strings.TrimSuffix(fileHeader.Filename, extension)
	nowFilename := fmt.Sprintf("%s_%d%s", filename, now.Nanosecond(), extension)
	fullFilepath := filepath.Join(FilesPath, nowFilename)

	f, openFileErr := os.OpenFile(fullFilepath, os.O_WRONLY|os.O_CREATE, 0666)

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
