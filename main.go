package main

import (
	"fmt"
	"gemini/config"
	"gemini/database"
	"gemini/router"
	"net/http"
)

func main() {

	config.Init()
	database.Init()
	router := router.Init()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.GetConfig().Port),
		Handler: router,
		// ReadTimeout: ,
		// WriteTimeout: ,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
