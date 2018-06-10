package types

import (
	"errors"
	"time"

	"github.com/oneiro-ndev/ndaumath/pkg/constants"
)

//go:generate msgp -tests=0

// A Timestamp is a single moment in time.
//
// It is monotonically increasing with the passage of time, and represents
// the number of microseconds since the epoch. It has no notion of leap time,
// time zones, or other complicating human factors.
// A timestamp can never be negative, but for mathematical simplicity we represent
// it with an int64. The total range of timestamps is almost 300,000 years.
type Timestamp int64

// A Duration is the difference between two Timestamps.
//
// It can be negative if the timestamps are out of order.
type Duration int64

// ParseTimestamp creates a timestamp from an ISO-3933 string
func ParseTimestamp(s string) (Timestamp, error) {
	ts, err := time.Parse(constants.TimestampFormat, s)
	if err != nil {
		return 0, err
	}
	return TimestampFrom(ts)
}

// TimestampFrom creates a Timestamp given a time.Time object
func TimestampFrom(t time.Time) (Timestamp, error) {
	// because this uses the standard library, it will overflow
	// some 290 years after the epoch
	//
	// TODO: implement this in a way which ensures its monotonic properties
	durationSinceEpoch := t.Sub(constants.Epoch)
	if durationSinceEpoch < 0 {
		return Timestamp(0), errors.New("date is before Epoch start")
	}
	return Timestamp(int64(durationSinceEpoch / time.Microsecond)), nil
}

// Compare implements comparison for Timestamp.
// (useful in sorting)
func (t Timestamp) Compare(o Timestamp) int {
	if t < o {
		return -1
	} else if t > o {
		return 1
	}
	return 0
}

// Since measures the Duration between two Timestamps.
// It will be positive when the argument is older, so present.Since(past) > 0
func (t Timestamp) Since(o Timestamp) Duration {
	return Duration(t - o)
}

// Add adds the supplied Duration to the given Timestamp
// If the result is negative, returns 0
// If the result overflows, returns MaxTimestamp
func (t Timestamp) Add(d Duration) Timestamp {
	ts := Timestamp(int64(t) + int64(d))
	if ts < constants.MinTimestamp {
		if d < 0 {
			return constants.MinTimestamp
		}
		return constants.MaxTimestamp
	}
	return ts
}

// Sub subtracts the supplied Duration from the given Timestamp
func (t Timestamp) Sub(d Duration) Timestamp {
	ts := Timestamp(int64(t) - int64(d))
	if ts < constants.MinTimestamp {
		if d > 0 {
			return constants.MinTimestamp
		}
		return constants.MaxTimestamp
	}
	return ts
}

func (t Timestamp) String() string {
	tt := constants.Epoch.Add(time.Duration(t) * time.Microsecond)
	return tt.Format(constants.TimestampFormat)
}

const (
	// Millisecond is a thousandth of a second
	Millisecond = 1
	// Second is the duration of 9 192 631 770 periods of the
	// radiation corresponding to the transition between the two
	// hyperfine levels of the ground state of the cesium 133 atom,
	// per the 13th CGPM (1967).
	Second = Millisecond * 1000
	// Day is exactly 86400 Seconds
	Day = Second * 86400
	// Year is exactly 365 days
	Year = Day * 365
)