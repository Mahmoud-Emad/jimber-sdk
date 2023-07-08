build:
	go build cmd/server.go

run: build
	./server

test: build
	go test ./...

clean:
	rm -rf ./server