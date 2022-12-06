package db

import (
	"github.com/istomin10593/bookstore_oauth-api/src/clients/cassandra"
	"github.com/istomin10593/bookstore_oauth-api/src/domain/access_token"
	"github.com/istomin10593/bookstore_oauth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}

}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	var result access_token.AccessToken

	if session := cassandra.GetSession(); session == nil {
		return nil, errors.NewInternalServerError("Error")
	}
	return &result, nil
}
