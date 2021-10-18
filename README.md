# Queue Process Portal
**tldr:** This is yet another process manager, but this time this is using queue messages to spawn processes.
## Required params:
- Redis address (default `127.0.0.1:6379`)
- Redis password (default `""`)
- Redis DB (default `0`)
- Redis List name of the list that runner will run `rpop` on
- Max number of PHP processes (default `1`)

## Flags
- `--redis.address`
- `--redis.password`
- `--redis.db`
- `--redis.list`
- `--workers.max`
- `--workers.command`
- `--version`

## Env vars
- `REDIS_HOST`
- `REDIS_PASSWORD`
- `REDIS_DB`
- `REDIS_LIST`
- `WORKERS_MAX`
- `WORKERS_COMMAND`