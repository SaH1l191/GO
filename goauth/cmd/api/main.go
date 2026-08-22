package cmd

import (
	"go-auth/internal/httpServer"
	"log"
	"net/http"
	"time"
)

func main() {
	router := httpServer.NewRouter()
	srv := &http.Server{
		Addr:    ":8080",
		Handler : router,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("Starting server on %s", srv.Addr)
	
	if err := srv.ListenAndServe(); err != nil  {
		if err == http.ErrServerClosed {
			log.Printf("server closed! ")
			return 
		} 
		log.Fatalf("Could not listen on %s: %v\n", srv.Addr, err)
	} 
}