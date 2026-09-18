# Deletion, Eviction, and Persistence

**Terminal session:**

```
> redis-cli
127.0.0.1:6379> config get maxmemory
1) "maxmemory"
2) "1073741824"
127.0.0.1:6379> config get maxmemory-policy
1) "maxmemory-policy"
2) "noeviction"
127.0.0.1:6379> config set maxmemory-policy allkeys-lru
OK
```

**Annotations (right side):**

- Check the maximum memory limit — when this threshold is reached, memory eviction is triggered, i.e. some keys get deleted.
- Check the memory eviction policy — the default is that once the threshold is reached, inserting another key simply throws an error.
- Set the memory eviction policy — generally LRU or LFU is used.

**Bullet points:**

- **LRU (Least Recently Used)** — based on a linked-list structure, where elements are ordered front to back by operation order. The most recently operated-on key is moved to the head of the list, so when memory eviction is needed, you only need to delete the element at the tail of the list.
- **LFU (Least Frequently Used)** — its basic assumption is that if data has been accessed many times in the past, it will likely be accessed more frequently in the future as well. So it evicts the keys with the lowest past access frequency. Redis uses a complex but efficient method to approximate LFU.
- LFU is somewhat more reasonable than LRU, but more complex to implement.
- **Expired-data deletion strategies:** periodic deletion and lazy deletion. In short, a key is not deleted the instant it expires.
- Redis can also start background threads to persist data (write it to disk).
