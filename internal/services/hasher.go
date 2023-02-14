package services

import (
	"time"

	"example.com/hasher/internal/domain"
	"example.com/hasher/pkg/cacher"
)

// Hasher service
type Hasher struct {
	cacher *cacher.MemoryCacher[domain.Hash]
}

// Dependency. Hasher repository
type HasherRepository interface {
	Generate() *domain.Hash
}

// Instantiate hasher service
func NewHasher(
	cacheInterval time.Duration,
	hasherRepository HasherRepository,
) *Hasher {
	return &Hasher{
		cacher.NewMemoryCacher(
			cacheInterval,
			hasherRepository.Generate,
		),
	}
}

// Get hash from cacher
func (c *Hasher) GetHash() *domain.Hash {
	return c.cacher.GetResult()
}
