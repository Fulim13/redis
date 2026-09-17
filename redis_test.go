package redis

import (
	"context"
	"log/slog"
	"testing"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       0,  // Redis creates DB 0-15 by default; this uses the default DB
		Password: "", // No password
	})
	// The connection is only established successfully if ping succeeds
	if err := client.Ping(context.Background()).Err(); err != nil {
		slog.Error("connect to redis failed", "error", err)
	} else {
		slog.Info("connect to redis")
	}
}

func TestStringValue(t *testing.T) {
	StringValue(context.Background(), client)
}

func TestStructValue(t *testing.T) {
	ctx := context.Background()
	stu := &Student{Id: 1, Name: "Fu Lim"}
	if err := WriteStudent2Redis(ctx, client, stu); err != nil {
		t.Fatal(err)
	}
	stu2 := GetStudentFromRedis(ctx, client, 1)
	if stu2 == nil {
		t.Fatal("student not found in redis")
	}
	if stu2.Id != stu.Id {
		t.Fail()
	}
	if stu2.Name != stu.Name {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	DeleteKey(context.Background(), client)
}

func TestScan(t *testing.T) {
	Scan(context.Background(), client)
}

func TestListValue(t *testing.T) {
	ListValue(context.Background(), client)
}

func TestSetgValue(t *testing.T) {
	SetValue(context.Background(), client)
}

func TestZSetValue(t *testing.T) {
	ZsetValue(context.Background(), client)
}

func TestHashTableValue(t *testing.T) {
	HashtableValue(context.Background(), client)
}

// go test -v -run='^TestStringValue$' -count=1
// go test -v -run='^TestStructValue$' -count=1
// go test -v -run='^TestDelete$' -count=1
// go test -v -run='^TestListValue$' -count=1
// go test -v -run='^TestSetgValue$' -count=1
// go test -v -run='^TestZSetValue$' -count=1
// go test -v -run='^TestHashTableValue$' -count=1
// go test -v -run='^TestScan$' -count=1
