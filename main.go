package main

import (
	"fmt"
	"gemini/config"
	"gemini/database"
	"gemini/router"
	"gemini/router/api"
	"net/http"
)

func main() {

	config.ReadConfig()

	database.Init()

	api.InitValidator()

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
