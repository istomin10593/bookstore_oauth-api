package db

import (
	"github.com/istomin10593/bookstore_oauth-api/src/domain/access_token"
	"github.com/istomin10593/bookstore_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}
