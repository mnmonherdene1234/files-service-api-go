package main

import (
	"FilesServiceAPI/api"
	"FilesServiceAPI/config"
)

func main() {
	config.Start()
	api.Start()
}
