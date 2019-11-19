# gqlgen & apollo federation

https://github.com/marwan-at-work/gqlgen/tree/federation

```shell script
$ npx prettier --write "./*/*.graphql"
$ goimports -w ./**/*.go
$ go generate ./...
```

## :eyes: Issues

* `extends Query` will make a panic
    * > Cannot extend type Query because it does not exist.
    * use `@extends` instead
