package resolvers

import "github.com/vikusku/go-gql-server/internal/orm"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	ORM *orm.ORM
}
