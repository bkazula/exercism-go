package gigasecond

import (
	"math"
	"time"
)

func AddGigasecond(t time.Time) time.Time {
	gigasecond := int64(math.Pow10(9))
	return t.Add(time.Second * time.Duration(gigasecond))
}
