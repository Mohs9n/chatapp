package internal

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Mohs9n/chat/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var query *db.Queries

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Getenv("DB_URL"))
	con, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	query = db.New(con)
}

// creates and returns the gin router
func NewRouter() (router *gin.Engine) {
	router = gin.Default()
	
	// Unprotected routes
	unprotected := router.Group("/")
	{
		unprotected.POST("/login", LoginHandler)
		unprotected.POST("/signup", RegisterHandler)
		unprotected.POST("/usersearch", userSearchByName)
	}

	// Protected routes
	protected := router.Group("/")
	protected.Use(authMiddleware())
	{
		protected.POST("/friend", sendFriendRequestHandler)
		protected.GET("/friend", getFriendRequests)
		protected.POST("/friend/accept", acceptFriendRequest)
	}

	return router
}