start server

```
go run server.go
```

if you edit graphql schema file, you need to generate graphql schema

```
go get github.com/99designs/gqlgen

go run github.com/99designs/gqlgen generate
```

```
add new ent schema

go run entgo.io/ent/cmd/ent new User
```

if you edit ent schema file, you need to generate ent schema

```
go run entgo.io/ent/cmd/ent generate ./ent/schema

OR
(recommended)
go run ent/entc.go
```

Build server
```
Windows
-----------------------
go build -o ydsapi.exe server.go

Linux
-----------------------
set GOOS=linux&& set GOARCH=amd64&& go build -o ydsapi server.go
```
