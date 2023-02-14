package domain

import (
	"time"

	"github.com/google/uuid"
)

// Hash model
type Hash struct {
	Value     uuid.UUID
	CreatedAt time.Time
}
