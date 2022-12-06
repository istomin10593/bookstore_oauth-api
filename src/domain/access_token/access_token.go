package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/istomin10593/bookstore_oauth-api/src/utils/crypto_utils"
	"github.com/istomin10593/bookstore_oauth-api/src/utils/errors"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before((time.Now().UTC()))
}

func (at *AccessToken) Generate() *errors.RestErr {
	hash, err := crypto_utils.HashedValue(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
	if err != nil {
		// logger.Error("error when trying to get hashed value", err)
		restErr := errors.NewInternalServerError("database error")
		return restErr
	}
	at.AccessToken = hash
	return nil

}
