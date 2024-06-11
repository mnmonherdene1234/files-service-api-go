package api

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc(Prefix, Handler)

	fmt.Printf("Server is running on port %d\n", Port)
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%d", Port),
			AllowCORS(http.HandlerFunc(Handler)),
		),
	)
}
