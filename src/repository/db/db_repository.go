package db

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/istomin10593/bookstore_oauth-api/src/clients/cassandra"
	"github.com/istomin10593/bookstore_oauth-api/src/domain/access_token"
	"github.com/istomin10593/bookstore_oauth-api/src/logger"
	"github.com/istomin10593/bookstore_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}

}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	var result access_token.AccessToken

	if getErr := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); getErr != nil {
		if getErr == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError(fmt.Sprintf("no access token found with given id=%s", id))
		}
		logger.Error("error when trying to get current id", getErr)
		return nil, errors.NewInternalServerError("database error")
	}

	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
	if saveErr := cassandra.GetSession().Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); saveErr != nil {
		logger.Error("error when trying to save access token", saveErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		logger.Error("error when trying to update current resource", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
