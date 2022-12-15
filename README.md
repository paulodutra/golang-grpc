# Golang grpc

The project that using golang 1.19. For learning about grpc.

1. For install dependencies: go mod tidy

2. For generate proto files (protocol buffer and grpc):

```bash
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

3. For install gRPC dependencies: go mod tidy

4. create database and table:

```
touch db.sqlite
sqlite3 db.sqlite
create table categories (id string, name string, description string);
```