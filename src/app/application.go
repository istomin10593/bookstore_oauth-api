package app

import (
	"github.com/gin-gonic/gin"
	"github.com/istomin10593/bookstore_oauth-api/src/http"
	"github.com/istomin10593/bookstore_oauth-api/src/repository/db"
	"github.com/istomin10593/bookstore_oauth-api/src/service/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token", atHandler.GetById)

	router.Run(":8080")
}
