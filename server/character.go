package server

import (
	"database/sql"
	"github.com/EliriaT/dnd-user-service/db"
	"github.com/EliriaT/dnd-user-service/server/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) createCharacter(ctx *gin.Context) {
	var uri dto.GetUserByIdRequest
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid or missing id parameter"})
		return
	}

	_, err := server.queries.GetUserByID(ctx, uri.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var req dto.CreateCharacterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	character, err := server.queries.CreateCharacter(ctx, db.CreateCharacterParams{
		UserID: uri.ID,
		Name:   req.Name,
		Traits: req.Traits,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, struct {
		ID int64 `json:"id"`
	}{
		ID: character.ID,
	})
}

func (server *Server) getCharactersByUserID(ctx *gin.Context) {
	var uri dto.GetUserByIdRequest
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid or missing id parameter"})
		return
	}

	_, err := server.queries.GetUserByID(ctx, uri.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	characters, err := server.queries.GetCharactersByUserID(ctx, uri.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, characters)
}
