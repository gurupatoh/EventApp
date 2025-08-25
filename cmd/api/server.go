package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)
func (app *application) server() error {	
	server := &http.Server{
	Addr:    fmt.Sprintf(":%d", app.port),
	Handler: app.routes(),
	IdleTimeout: 30 * time.Second,

}
log.Printf("Starting server on port: %d", app.port)

return server.ListenAndServe()
}
