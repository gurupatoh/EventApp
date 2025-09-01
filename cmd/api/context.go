package main

import (
	"EventApp/internal/database"
	"github.com/gin-gonic/gin"
)

func (app *application) GetUserFromContext(C*gin.Context) *database.User {
	contextUser, exists := C.Get("user")
	if !exists {
		return &database.User{}
	}
	user, ok := contextUser.(*database.User)
	if !ok {
		return &database.User{}
	}
	return user
}