package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.POST("/token", authorize)
	r.Run()
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

type authCode struct {
	code string
}

func authorize(c *gin.Context) {
	var newAuthCode authCode
	if err := c.BindJSON(&newAuthCode); err != nil {
		fmt.Println(err)
		return
	}

	myConfig := &oauth2.Config{
		ClientID:     "<client_id>",
		ClientSecret: "<client_secret>",
		Endpoint:     google.Endpoint,
		Scopes:       []string{"scopes"},
		RedirectURL:  "<redirect_uri>",
	}

	myToken := token(myConfig, newAuthCode.code)

	c.JSON(200, gin.H{
		"token": myToken,
	})
}

func token(config *oauth2.Config, code string) *oauth2.Token {
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		fmt.Printf("unable to retrieve token from web: %v", err)
	}

	return token
}
