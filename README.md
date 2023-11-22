# gqlgen & apollo federation

this repository is port of https://github.com/apollographql/federation-demo

with https://gqlgen.com/


## How to use

```shell script
$ docker compose up --build
$ open http://localhost:4000/
```

or

```shell script
$ go run ./accounts/server/server.go
$ go run ./reviews/server/server.go
$ go run ./products/server/server.go
$ go run ./inventory/server/server.go
$ (cd gateway && npm ci && npm run start)
$ open http://localhost:4000/
```

## Development

```shell script
$ npx prettier --write "./*/*.{graphql,js}"
$ goimports -w ./**/*.go
$ go generate ./...
```
