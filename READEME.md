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

# Reference

https://youtu.be/xaA_Gb0L_kQ?si=dOjYhKbo1MmleUiN
https://youtu.be/Fd6KMDjBY34?si=SJVZxi179GQ9hGKk
https://youtu.be/dx0LdO0GbbQ?si=g3v0PD7GWtYC9xyi
