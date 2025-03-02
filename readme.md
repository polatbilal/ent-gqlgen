start server

```
go run main.go
```

if you edit graphql schema file, you need to generate graphql schema

```
go get github.com/99designs/gqlgen

go run github.com/99designs/gqlgen generate
```

if you edit ent schema file, you need to generate ent schema

```
go run entgo.io/ent/cmd/ent generate ./api-core/ent/schema

OR

go run api-core/ent/entc.go (recommended)
```
