package server

import (
	"log"

	"GolangAPI/server/routes"

	"github.com/gin-gonic/gin"
)

type server struct {
	port   string
	server *gin.Engine
}

func newServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Print("server is runnig at port: ", s.port)
	log.Fatal(Router.run(":" + s.port))
}
