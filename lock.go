package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// TryLock tries to acquire the distributed lock. It returns true on success and false on failure.
func TryLock(ctx context.Context, rc *redis.Client, key string, expire time.Duration) bool {
	cmd := rc.SetNX(ctx, key, "any value", expire) // SetNX returns true if the key does not exist, writes the key and sets its TTL
	if cmd.Err() != nil {
		return false
	} else {
		return cmd.Val()
	}
}

// ReleaseLock releases the distributed lock
func ReleaseLock(ctx context.Context, rc *redis.Client, key string) {
	rc.Del(ctx, key)
}
