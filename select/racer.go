package racer

import (
	"net/http"
	"time"
)

func Racer(url1, url2 string) (winner string) {
	aDuration := _measureResponseTime(url1)
	bDuration := _measureResponseTime(url2)

	if aDuration < bDuration {
		return url1
	}

	return url2
}

func _measureResponseTime(url string) time.Duration {
	start := time.Now()
	_, _ = http.Get(url)
	duration := time.Since(start)
	return duration
}
