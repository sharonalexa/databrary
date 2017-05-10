package routes

import (
	"gopkg.in/throttled/throttled.v2"
	"gopkg.in/throttled/throttled.v2/store/memstore"
)

func NewRateLimiter() (throttled.HTTPRateLimiter, error) {

	store, err := memstore.New(65536)
	if err != nil {
		return nil, err
	}

	quota := throttled.RateQuota{throttled.PerSec(2), 5}
	rateLimiter, err := throttled.NewGCRARateLimiter(store, quota)
	if err != nil {
		return nil, err
	}

	return throttled.HTTPRateLimiter{
		RateLimiter: rateLimiter,
		VaryBy:      &throttled.VaryBy{RemoteAddr: true},
	}, nil
}
