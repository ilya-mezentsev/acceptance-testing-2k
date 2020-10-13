package timers

import "time"

func ConnectionCacheCleanTimout() time.Duration {
	return 20 * time.Minute
}

func ConnectionCacheLifetime() time.Duration {
	return 30 * time.Minute
}

func DeletedAccountHashesCleanTimeout() time.Duration {
	return time.Hour
}

func DeletedAccountHashesCacheLifetime() time.Duration {
	return 12 * time.Hour
}

func CheckExpiredAccountTimeout() time.Duration {
	return time.Hour
}
