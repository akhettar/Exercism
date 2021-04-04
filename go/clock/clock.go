package clock

import (
	"fmt"
)

// Clock type
type Clock struct {
	m int
}

// New creates an instance of the Clock type
func New(h, m int) Clock {

	// calculate total minutes including the hours
	m += h * 60

	// mintues must be less than a day
	m %= 24 * 60

	// rework out the number of minutes from 24 hour base backwards
	if m < 0 {
		m += 24 * 60
	}
	return Clock{m}
}

// String implment the Stringer interface
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.m/60, c.m%60)
}

// Add minutes to the clock
func (c Clock) Add(m int) Clock {
	return New(0, c.m+m)
}

// Subtract minutes from the clock
func (c Clock) Subtract(m int) Clock {
	return New(0, c.m-m)
}
