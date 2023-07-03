package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/session"
)

var (
	key   = []byte("super-secret-key")
	store = session.NewCookieStore(key)
)

func GetCurrentUser(c *gin.Context) {
	session, _ := store.Get(c.Request, "session-name")
	userID, ok := session.Values["user_id"].(unit)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// TODO: Get user from database using userID

	c.JSON(http.StatusOK, gin.H{"user_id": userID})
}

func Login(c *gin.Context) {
	// TODO Authenticate user

	session, _ := store.Get(c.Request, "session-name")
	session.Values["user_id"] = user.ID // replace user.ID with the authenticated user's ID
	session.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func Logout(c *gin.Context) {
	session, _ := store.Get(c.Request, "session-name")
	delete(session.Values, "user_id")
	session.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, gin.H{"status": "you are logged out"})
}
