package transformation

import (
	"errors"
	"math/rand"
	"strconv"

	gql "github.com/vikusku/go-gql-server/internal/gql/models"
	dbm "github.com/vikusku/go-gql-server/internal/orm/models"
)


var r = rand.New(rand.NewSource(99))

// DBUserToGQLUser transforms [user] db input to gql type
func DBUserToGQLUser(i *dbm.User) (o *gql.User, err error) {
	o = &gql.User{
		ID:          strconv.FormatUint(i.ID, 10),
		Email:       i.Email,
		UserID:      i.UserID,
		Name:        i.Name,
		FirstName:   i.FirstName,
		LastName:    i.LastName,
		NickName:    i.NickName,
		Description: i.Description,
		Location:    i.Location,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
	}
	return o, err
}

// GQLInputUserToDBUser transforms [user] gql input to db model
func GQLInputUserToDBUser(i *gql.UserInput, update bool, ids ...string) (o *dbm.User, err error) {
	o = &dbm.User{
		UserID:      i.UserID,
		Name:        i.Name,
		FirstName:   i.FirstName,
		LastName:    i.LastName,
		NickName:    i.NickName,
		Description: i.Description,
		Location:    i.Location,
	}
	if i.Email == nil && !update {
		return nil, errors.New("Field [email] is required")
	}
	if i.Email != nil {
		o.Email = *i.Email
	}
	if len(ids) > 0 {
		o.ID = r.Uint64()
	}
	return o, err
}