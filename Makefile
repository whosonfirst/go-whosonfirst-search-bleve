cli:
	go build -mod vendor -o bin/index cmd/index/main.go
	go build -mod vendor -o bin/query cmd/query/main.go
