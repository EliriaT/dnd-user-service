package server

import (
	"github.com/EliriaT/dnd-user-service/db"
	"github.com/EliriaT/dnd-user-service/server/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) createUser(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	user, err := server.queries.CreateUser(ctx, db.CreateUserParams{
		Email:    req.Email,
		Password: hashedPassword,
		Username: req.Username,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, struct {
		Id       int64  `json:"id"`
		Username string `json:"username"`
	}{
		Id:       user.ID,
		Username: user.Username,
	})
}
