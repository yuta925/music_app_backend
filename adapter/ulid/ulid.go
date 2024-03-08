package ulid

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid"
	"music-app/usecase/port"
)

func GenerateULID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.Reader, 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

type ULID struct{}

func NewULID() port.ULID {
	return &ULID{}
}

func (u *ULID) GenerateID() string {
	return GenerateULID()
}
