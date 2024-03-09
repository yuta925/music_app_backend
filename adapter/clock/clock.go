package clock

import (
	"time"
	"music-app/usecase/port"
)

type Clock struct{}

func New() port.Clock {
	return Clock{}
}

func (c Clock) Now() time.Time {
	return time.Now()
}
