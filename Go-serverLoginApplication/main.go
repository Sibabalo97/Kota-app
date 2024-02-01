package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("145821")

// User represents the user information for authentication
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JWTClaims represents the claims used in the JWT
type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	router := gin.Default()

	// Registration endpoint
	router.POST("/api/register", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Perform user registration logic here
		// Save user to the database or any other necessary operations

		// Create JWT token
		token, err := createJWTToken(user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JWT token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	// Login endpoint
	router.POST("/api/login", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Perform user login logic here
		// Check username and password against the database or any other authentication mechanism

		// Create JWT token
		token, err := createJWTToken(user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JWT token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	// Run the server
	port := 5000
	address := fmt.Sprintf(":%d", port)
	log.Printf("Server is running on http://localhost:5000", port)
	log.Fatal(router.Run(address))
}

// createJWTToken creates a JWT token with the given username
func createJWTToken(username string) (string, error) {
	claims := JWTClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
