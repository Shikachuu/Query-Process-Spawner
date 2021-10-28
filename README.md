# Queue Process Portal
[![Go Report Card](https://goreportcard.com/badge/github.com/Shikachuu/php-process-redis-list)](https://goreportcard.com/report/github.com/Shikachuu/php-process-redis-list)

**tldr:** This is yet another process manager, but this time this is using queue messages to spawn processes.

## Root flags
| Flag              | Description                                                     | Default | Example usage           |
|-------------------|-----------------------------------------------------------------|---------|-------------------------|
| `--max-workers`   | Maximum number of running processes                             | `1`     | `qpp --max-workers 2`   |
| `-c`, `--command` | Command to handle message, comma separated list with the params | `""`    | `qpp -c php,worker.php` |

## Queues
### Redis list
Command: `qpp redis-list`

| Flag         | Description               | Default          | Example usage                                    |
|--------------|---------------------------|------------------|--------------------------------------------------|
| `--host`     | Database host name        | `127.0.0.1:6379` | `qpp redis-list --host my-super-redis-host:6379` |
| `--password` | Database password         | `""`             | `qpp redis-list --password SuperSecret99+!`      |
| `--db`       | Number of redis database  | `0`              | `qpp redis-list --db 1`                          |
| `--list`     | Redis list to `rpop` from | `""`             | `qpp redis-list --list test2`                    |

## How-to build
You must have at least go version 1.11 and GNU Make to build this.

Run `make build` to build a dev version. (Builds only `linux/amd64`)

Run `make dist` to build "release" version. (It cross-compiles by default to `linux/arm64` and `linux/amd64`)

