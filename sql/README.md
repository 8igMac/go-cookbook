# How to Access Relational Databases with Go?

Using PostgreSQL as an example.

## Packages Used

- [sql](https://pkg.go.dev/database/sql): Generic SQL interface provided
  by the Go standard library. Must be used in conjunction with a
  database driver.
- [pq](https://github.com/lib/pq): Pure Go PostgreSQL driver for
  database/sql.

## Prerequisite

- Install Docker.

## Quick Start

- Start a Postgres docker instance with docker-compose.

```sh
cd sql
docker compose up -d
```

- Execute the program.

```sh
go run main.go
```

- Tear down the Postgres docker instance.

```sh
docker compose down
```
