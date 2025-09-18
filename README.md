# Quiz

## Development

### Build

```sh
go build ./cmd/quiz_server
```

## TODO

* User registration service
* REST service to create/update/delete quiz

## Dependency

### Golang

### Goosee ??

### PostgreSQL

* Download

```sh
docker pull postgres:13-trixie
```

* Start

```sh
docker run --name quiz_db -e POSTGRES_USER=quiz_db_user -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=quiz_db -p 5432:5432 -d postgres
```

### Echo?
