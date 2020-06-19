package app

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vadim-dmitriev/sittme/state"
)

type streamNotFoundError struct {
	uuid uuid.UUID
}

func (e streamNotFoundError) Error() string {
	return fmt.Sprintf("stream '%s' not found", e.uuid)
}

type canNotChangeStateError struct {
	uuid         uuid.UUID
	currentState state.Stater
	newState     state.Stater
}

func (e canNotChangeStateError) Error() string {
	return fmt.Sprintf("can not change '%s' state to '%s' state at '%s' stream",
		e.currentState,
		e.newState,
		e.uuid,
	)
}
