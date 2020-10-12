package timers

import "time"

func ConnectionCacheCleanTimout() time.Duration {
	return 20 * time.Minute
}

func ConnectionCacheLifetime() time.Duration {
	return 30 * time.Minute
}

func CheckExpiredAccountTimeout() time.Duration {
	return time.Hour
}
