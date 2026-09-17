package redis

import (
	"context"
	"testing"
)

func TestPubSub(t *testing.T) {
	PubSub(context.Background(), client)
}

// go test -v . -run='^TestPubSub$' -count=1
