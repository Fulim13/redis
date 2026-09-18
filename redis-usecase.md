# Common Use Cases

- **General principle:** Redis is an order of magnitude faster than MySQL, but less reliable than MySQL. So data that needs high-frequency reads/writes, has a short lifecycle, and isn't especially critical to the user is well suited to being stored in Redis.

- **Counters.** `Incr(ctx context.Context, key string)` increments the count for the corresponding key by 1 — for example, video play counts, or product inventory in flash-sale scenarios. `INCRBY` adds an arbitrary value, which can be negative.

- **Caching.** Frequently accessed MySQL data can be put into Redis, with the key corresponding to the id and the value being a JSON string. This reduces pressure on MySQL and improves API response speed.

- **Session caching.** The SessionID is used to mark a user as successfully logged in; the login and subsequent operations may hit different servers, so the SessionID needs to be stored in a distributed cache. Search/recommendation result lists are stored in the cache and read from it when paging.

- **Distributed locks.** In a distributed system, a scheduled task only needs to be executed by one server — whoever grabs the lock executes it, and releases the lock before the next cycle arrives.

- **Publish/subscribe functionality.** A small volume of event notifications can be implemented with Redis; large-volume message delivery is better suited to Kafka.
