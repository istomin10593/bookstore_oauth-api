package access_token

import (
	"github.com/istomin10593/bookstore_oauth-api/src/domain/access_token"
	"github.com/istomin10593/bookstore_oauth-api/src/repository/db"
	"github.com/istomin10593/bookstore_oauth-api/src/utils/errors"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type service struct {
	dbRepo db.DbRepository
}

func NewService(dbRepo db.DbRepository) Service {
	return &service{
		dbRepo: dbRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, nil
}
