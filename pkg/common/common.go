package common

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// ULID Generator
func GenerateULID() string {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
