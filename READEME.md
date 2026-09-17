# To Install Redis

```sh
brew install redis
```

# Start Service

```sh
brew services list
brew services start redis
```

# Test Redis

```sh
redis-cli
set name fulim
get name
set age 24
# The returned age is string
get age
```
