package server

import (
	"fmt"
	"gorm.io/gorm"
	"ju-go-server/config"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	DB         *gorm.DB
}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) Run(port string) error {
	router := InitRouter(server)
	config.SetupRoutes(router)

	server.httpServer = &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := server.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	fmt.Println("Success!")
	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	return nil
}

func (server *Server) Stop() error {
	err := server.httpServer.Close()
	if err != nil {
		return err
	}
	return nil
}
