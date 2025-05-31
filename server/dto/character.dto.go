package dto

import "encoding/json"

type CreateCharacterRequest struct {
	Name   string          `json:"name" binding:"required"`
	Traits json.RawMessage `json:"traits" binding:"required"`
}
