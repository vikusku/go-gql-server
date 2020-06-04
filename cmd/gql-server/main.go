package main

import (
	"github.com/vikusku/go-gql-server/internal/orm"
	"github.com/vikusku/go-gql-server/pkg/server"
	"log"
)

func main() {
	// Create a new ORM instance to send it to our
	orm, err := orm.Factory()
	if err != nil {
		log.Panic(err)
	}

	server.Run(orm)
}