package server

import (
	"github.com/vikusku/go-gql-server/pkg/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vikusku/go-gql-server/internal/handlers"
)

var HOST, PORT string

func init() {
	HOST = utils.MustGet("GQL_SERVER_HOST")
	PORT = utils.MustGet("GQL_SERVER_PORT")
}

// Run web server
func Run() {
	r := gin.Default()
	// Setup routes
	r.GET("/ping", handlers.Ping())
	log.Println("Running @ http://" + HOST + ":" + PORT )
	log.Fatalln(r.Run(HOST + ":" + PORT))
}