package slidingwindow

import "github.com/amitkatyal/rate-limiter/pkg/datastore"

type WindowStatus struct {
	Allow bool
}

type MinuteWindow struct {
	size int
}

type SecondWindow struct {
	size int
}

// SlidingWindowLimiter rate limiting interface
type SlidingWindowLimiter interface {
	Allow() bool
	AllowWithStatus() WindowStatus
}

type TimeWindow interface {
	Current() int
	Previous (int, int)
}


// SlidingWindowLimiterImpl implements sliding window rate limiting algorithm
type SlidingWindowLimiterImpl struct {
	DataStore datastore.DataStore
	RequestLimit uint32
	RequestRate TimeWindow
}

func NewPerMinuteRequestRate() *MinuteWindow {
	return &MinuteWindow{size: 1}
}

func NewPerSecondRequestRate() *SecondWindow {
	return &SecondWindow{size: 1}
}


func (rate *MinuteWindow) current() int {
	return 0
}

func NewSlidingWindow(dataStore datastore.DataStore, requestLimit uint32, requestRate TimeWindow)  SlidingWindowLimiter {
	return &SlidingWindowLimiterImpl{
		DataStore: dataStore,
		RequestLimit: requestLimit,
		RequestRate: requestRate,
	}
}

// Allow checks if request is allowed
func (sw *SlidingWindowLimiterImpl) Allow() bool {
	sw.RequestRate.Current()
	return true
}

// AllowWithStatus checks if request is allowed. If not allowed, return the window status
func (sw *SlidingWindowLimiterImpl) AllowWithStatus() WindowStatus {
	return WindowStatus{}
}