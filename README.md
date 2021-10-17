# Yet Another PHP process manager, but this time for redis lists

## Required params:
- Redis address (default `127.0.0.1:6379`)
- Redis password (default `""`)
- Redis DB (default `0`)
- Max number of PHP processes (default `1`)

## Flags
- `--redis.address`
- `--redis.password`
- `--redis.db`
- `--workers.max`
- `--version`

## Env vars
- `REDIS_HOST`
- `REDIS_PASSWORD`
- `REDIS_DB`
- `WORKERS_MAX`