package clock

import "fmt"

const testVersion = 4

type Clock struct {
	m int
}

func New(hour, minute int) Clock {
	c := Clock{(minute + hour*60) % 1440}

	return c.fixedHours()
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.m/60, c.m%60)
}

func (c Clock) Add(minutes int) Clock {
	c.m += minutes
	c.m %= 1440

	return c.fixedHours()
}

func (c Clock) fixedHours() Clock {
	if c.m < 0 {
		c.m += 1440
	}

	return c
}
