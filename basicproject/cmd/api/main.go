package main
import (
	"fmt"
	"log"
	"notes-api/internal/config"
	"notes-api/internal/db"
	"notes-api/internal/server"
)
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Config Error !", err)
	}
	client, database, err := db.Connect(cfg)
	if err != nil {
		log.Fatal("Database Connection Error !")
	}
	defer func() {
		if err := db.Disconnect(client); err != nil {
			log.Printf("Database Disconnection Error !", err)
		}
	}()
	router := server.NewRouter(database)
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	fmt.Printf("Starting server on port %s\n", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal("Server Error !", err)
	}
}
