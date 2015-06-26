package main

import (
	"github.com/cyarie/cyarie.com/backend/settings"
	"github.com/cyarie/cyarie.com/backend/router"
	"log"
	"net/http"
)

func main() {
	context := &settings.AppContext{}
	router := router.ApiRouter(context)
	log.Fatal(http.ListenAndServe(":8080", router))
}
