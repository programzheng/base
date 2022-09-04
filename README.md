# base
This project is build ecommerce an api server
## run
### go
```shell
go run .
```
### docker-compose
```shell
docker-compose up -d api
```
#### Makefile
```shell
make up
```
## generate *.pb.go
```shell
cd internal/grpc/proto && /
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto
```
## feature
- [x] admin user jwt authorization api
- [x] file upload and file response api
- [ ] admin manage post
- [ ] grpc play games
