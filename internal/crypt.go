package internal

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	// Generate a salt with a cost of 10
	salt, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	// Generate the hash with the salt
	hash := string(salt)

	return hash, nil
}

func verifyPassword(password string, hash string) bool {
	// Convert the hash string to a byte slice
	hashBytes := []byte(hash)

	// Compare the password with the hash
	err := bcrypt.CompareHashAndPassword(hashBytes, []byte(password))
	return err == nil
}

// jwt
func createToken(userId int) (string, error) {
	// Create a new token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	secret := []byte("mysecretkey")
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key
		return []byte("mysecretkey"), nil
	})
	if err != nil {
		return nil, err
	}

	// Verify the token and return the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
