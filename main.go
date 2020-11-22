package main

import (
	"fmt"
	"gemini/config"
	"gemini/database"
	"gemini/router"
	"net/http"
)

func main() {

	config.ReadConfig()

	database.Init()

	router := router.Init()
	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.CONFIG.Port),
		Handler: router,
		// ReadTimeout: ,
		// WriteTimeout: ,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
