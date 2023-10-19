package internal

import (
	"github.com/Mohs9n/chat/internal/db"
	"github.com/gin-gonic/gin"
)

type UserSearch struct {
	Username string `json:"username"`
}

func userSearchByName(c *gin.Context) {
	var name UserSearch
	err:= c.BindJSON(&name)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid request JSON"})
		return
	}

	users, err := query.SearchUsers(c, "%"+name.Username+"%")
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, users)
}

func sendFriendRequestHandler(c *gin.Context) {
	// Get the current user ID from the context
	userID, ok := c.Get("userId")
	if !ok {
		c.JSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	// Get the friend's user ID from the request body
	var friendID struct {
		ID int `json:"id"`
	}
	if err := c.BindJSON(&friendID); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request JSON"+ err.Error()})
		return
	}

	// Send the friend request
	req := db.CreateFriendRequestParams{
		SenderID: int32(userID.(int)),
		ReceiverID: int32(friendID.ID),
	}
	if err := query.CreateFriendRequest(c, req); err != nil {
		c.JSON(500, gin.H{"error": "Friend Id not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Friend request sent"})
}

func getFriendRequests(c *gin.Context) {
	userID, ok := c.Get("userId")
	if !ok {
		c.JSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	fRequests, err := query.GetFriendRequests(c, int32(userID.(int)))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if fRequests==nil {
		c.JSON(200, gin.H{"message": "No friend requests"})
		return
	}

	c.JSON(200, fRequests)
}

func acceptFriendRequest(c *gin.Context) {
	userID, ok := c.Get("userId")
	if !ok {
		c.JSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	var friendID struct {
		ID int `json:"id"`
	}
	if err := c.BindJSON(&friendID); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request JSON"+ err.Error()})
		return
	}
	
	req := db.AcceptFriendRequestParams{
		SenderID: int32(friendID.ID),
		ReceiverID: int32(userID.(int)),
	}
	if err := query.AcceptFriendRequest(c, req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Friend request accepted"})
}