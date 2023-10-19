package internal

import (
	"net/http"

	"github.com/Mohs9n/chat/internal/db"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler handles the login request
func LoginHandler(c *gin.Context) {
	// Get the JSON body and decode into credentials
	var credentials credentials
	err := c.BindJSON(&credentials)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}

	// Get the expected password from our in memory map
	if bool, err := query.UserExists(c, credentials.Username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid request"})
		return
	} else if !bool {
		c.JSON(400, gin.H{"message": "User does not exist"})
		return
	}

	user, err := query.GetUserByUsername(c, credentials.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid request"})
		return
	}

	// Check if the password exists
	if !verifyPassword(credentials.Password, user.Passwordhash) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	// Respond with the token
	token, err := createToken(int(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// RegisterHandler handles the register request
func RegisterHandler(c *gin.Context) {
	// Get the JSON body and decode into credentials
	var NewUser User
	err := c.BindJSON(&NewUser)
	if err != nil {
		c.JSON(400, gin.H{"message": "Could not decode request"})
		return
	}

	if bool, err := query.UserExists(c, NewUser.Username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error checking if user exists"})
		return
	} else if bool {
		c.JSON(400, gin.H{"message": "User already exists"})
		return
	}

	hashedPassword, err := hashPassword(NewUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error hashing password"})
		return
	}

	userParams := db.CreateUserParams{
		Username:     NewUser.Username,
		Passwordhash: hashedPassword,
		Firstname:    NewUser.FirstName,
		Lastname:     NewUser.LastName,
	}
	user, err := query.CreateUser(c, userParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating user"})
		return
	}

	token, err := createToken(int(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created", "token": token})
}