package api

import (
	"FilesServiceAPI/config"
	"fmt"
	"log"
	"net/http"
)

func Start() {
	port := config.Get("PORT")
	prefix := config.Get("PREFIX")

	http.HandleFunc(prefix, Handler)

	fmt.Printf("Listening on %s\n", port)
	fmt.Printf("Prefix: %s\n", prefix)
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%s", port),
			AllowCORS(http.HandlerFunc(Handler)),
		),
	)
}
