package server

import (
	"github.com/vikusku/go-gql-server/internal/orm"
	"github.com/vikusku/go-gql-server/pkg/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vikusku/go-gql-server/internal/handlers"
)

var host, port, gqlPath, gqlPgPath string
var isPgEnabled bool

func init() {
	host = utils.MustGet("GQL_SERVER_HOST")
	port = utils.MustGet("GQL_SERVER_PORT")
	gqlPath = utils.MustGet("GQL_SERVER_GRAPHQL_PATH")
	gqlPgPath = utils.MustGet("GQL_SERVER_GRAPHQL_PLAYGROUND_PATH")
	isPgEnabled = utils.MustGetBool("GQL_SERVER_GRAPHQL_PLAYGROUND_ENABLED")
}

// Run web server
func Run(orm *orm.ORM) {
	log.Println("GORM_CONNECTION_DSN: ", utils.MustGet("GORM_CONNECTION_DSN"))

	endpoint := "http://" + host + ":" + port

	r := gin.Default()
	// Setup routes
	r.GET("/ping", handlers.Ping())

	if isPgEnabled {
		r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		log.Println("GraphQL Playground @ " + endpoint + gqlPgPath)
	}

	r.POST(gqlPath, handlers.GraphqlHandler(orm))
	log.Println("GraphQL @ " + endpoint + gqlPath)

	log.Println("Running @ http://" + host + ":" + port )
	log.Fatalln(r.Run(host + ":" + port))
}