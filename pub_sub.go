package redis

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type CS string

func publish(ctx context.Context, client *redis.Client, channel string, message any) {
	cmd := client.Publish(ctx, channel, message)
	if cmd.Err() == nil {
		n := cmd.Val() // The number of subscribers
		fmt.Printf("%s published a message to channel %s, which currently has %d subscriber(s)\n", ctx.Value(CS("publisher_name")), channel, n)
	} else {
		fmt.Printf("%s failed to publish a message to channel %s: %v\n", ctx.Value(CS("publisher_name")), channel, cmd.Err())
	}
}

func subscribe(ctx context.Context, client *redis.Client, channels []string) {
	ps := client.Subscribe(ctx, channels...)
	defer ps.Close()

	// Loop forever, reading each message as soon as it arrives on the channel
	for {
		if msg, err := ps.ReceiveMessage(ctx); err != nil {
			// The context being cancelled is the normal way this loop ends
			if ctx.Err() == nil {
				fmt.Println(err)
			}
			break
		} else {
			fmt.Printf("%s received a message from channel %s: %s\n", ctx.Value(CS("subscriber_name")), msg.Channel, msg.Payload)
		}
	}
}

func PubSub(ctx context.Context, client *redis.Client) {
	// Cancelling this context stops every subscriber goroutine before the function returns
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ctx1 := context.WithValue(ctx, CS("publisher_name"), "publisher1")
	ctx2 := context.WithValue(ctx, CS("publisher_name"), "publisher2")
	channel1 := "channel1"
	channel2 := "channel2"

	// Start the first batch of subscribers.
	// A subscriber has to be started first; it cannot receive messages that were sent to the channel before that.
	ctx3 := context.WithValue(ctx, CS("subscriber_name"), "subscriber3")
	ctx4 := context.WithValue(ctx, CS("subscriber_name"), "subscriber4")

	go subscribe(ctx3, client, []string{channel1})
	go subscribe(ctx4, client, []string{channel2})
	time.Sleep(1 * time.Second)

	go publish(ctx1, client, channel1, "The white sun sets behind the mountains")
	go publish(ctx2, client, channel1, "The Yellow River flows into the sea")
	time.Sleep(1 * time.Second)
	fmt.Println(strings.Repeat("-", 50))

	// Start the second batch of subscribers
	ctx5 := context.WithValue(ctx, CS("subscriber_name"), "subscriber5")
	go subscribe(ctx5, client, []string{channel1, channel2})
	time.Sleep(1 * time.Second)

	go publish(ctx1, client, channel2, "To see a thousand miles further")
	go publish(ctx2, client, channel2, "Climb one more storey of the tower")
	time.Sleep(1 * time.Second)
}
