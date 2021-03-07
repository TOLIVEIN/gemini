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
		Addr:           fmt.Sprintf(":%s", config.GetConfig().Port),
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
		// ReadTimeout: ,
		// WriteTimeout: ,

	}
	// s.ListenAndServe()

	s.ListenAndServeTLS("config/fullchain.pem", "config/privkey.pem")
}
