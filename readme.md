start server

```
go run server.go
```

if you edit graphql schema file, you need to generate graphql schema

```
go get github.com/99designs/gqlgen

go run github.com/99designs/gqlgen generate
```

if you edit ent schema file, you need to generate ent schema

```
go run entgo.io/ent/cmd/ent generate ./ent/schema

OR

go run ent/entc.go (recommended)
```

Build server
```
go build -o ydsapi.exe ydsapi.go