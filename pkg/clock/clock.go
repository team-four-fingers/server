package clock

import "time"

type Clock interface {
	Now() time.Time
}

type clock struct {
	now func() time.Time
}

func (c *clock) Now() time.Time {
	return c.now()
}

func Default() Clock {
	return &clock{
		now: time.Now,
	}
}
