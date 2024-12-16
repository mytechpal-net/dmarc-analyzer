package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	// Initialize router
	router := newRouter()

	router.GET("/ping", ping)
	router.POST("/token", authorize)
	router.Run()
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

type authCode struct {
	Code string `json:"code"`
}

func authorize(c *gin.Context) {
	var newAuthCode authCode

	if err := c.BindJSON(&newAuthCode); err != nil {
		fmt.Printf("could not bind, %v", err)
		return
	}

	b, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, "scopes")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	myToken := token(config, newAuthCode.Code)

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

// Create gin router
func newRouter() *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies(nil)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:3000"},
		AllowMethods:     []string{"OPTIONS", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	return router
}
