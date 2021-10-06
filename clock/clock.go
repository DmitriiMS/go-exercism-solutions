package clock

import (
	"fmt"
)

type Clock struct {
	minutes int
}

func New(h, m int) Clock {
	min := (h*60 + m) % (60 * 24)
	if min < 0 {
		min += 60 * 24
	}
	c := Clock{min}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

func (c Clock) Add(min int) Clock {
	return New(0, c.minutes+min)
}

func (c Clock) Subtract(min int) Clock {
	return New(0, c.minutes-min)
}
