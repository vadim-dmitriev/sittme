package main

import (
	"github.com/google/uuid"
)

// generateUUID генерирует новый UUIDv4
func generateUUID() uuid.UUID {
	uuid, _ := uuid.NewRandom()

	return uuid
}
