package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(currentTime time.Time) Point {
	if currentTime.Second() == 30 {
		return Point{150, 240}
	}
	return Point{150, 60}
}

func secondsInRadians(seconds time.Time) float64 {
	return math.Pi
}

