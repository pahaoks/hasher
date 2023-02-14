package services

import (
	"testing"
	"time"

	"example.com/hasher/internal/repositories"
)

func TestIntervalUpdate(t *testing.T) {
	repo := repositories.NewHasher()
	service := NewHasher(
		time.Second*1,
		repo,
	)

	hash := service.GetHash()

	time.Sleep(time.Second * 2)

	hash2 := service.GetHash()

	if hash.Value == hash2.Value {
		t.Fatal("hash wasn't updated")
	}
}
