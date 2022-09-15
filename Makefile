
.PHONY: all test clean build

all: clean test build

build:
	go build -o bin/uppercase cmd/uppercase/main.go
	chmod +x bin/uppercase
	go build -o bin/lettercount cmd/lettercount/main.go
	chmod +x bin/lettercount
	go build -o bin/pages cmd/pages/main.go
	chmod +x bin/pages
	go build -o bin/timeout cmd/timeout/main.go
	chmod +x bin/timeout
	go build -o bin/panic cmd/panic/main.go
	chmod +x bin/panic
	go build -o bin/error cmd/error/main.go
	chmod +x bin/error
	go build -o bin/api cmd/api/main.go
	chmod +x bin/api

clean:
	go clean -cache -testcache -modcache
	rm -rf bin/