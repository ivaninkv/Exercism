package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	minutes int
}

func normalize(minutes int) int {
	minutes = minutes % 1440
	if minutes < 0 {
		minutes += 1440
	}
	return minutes
}

func New(h, m int) Clock {
	totalMinutes := h*60 + m
	return Clock{normalize(totalMinutes)}
}

func (c Clock) Add(m int) Clock {
	return Clock{normalize(c.minutes + m)}
}

func (c Clock) Subtract(m int) Clock {
	return Clock{normalize(c.minutes - m)}
}

func (c Clock) String() string {
	hours := c.minutes / 60
	minutes := c.minutes % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
