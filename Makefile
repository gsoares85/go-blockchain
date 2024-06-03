build:
	go mod tidy
	go build -o ./bin/go-blockchain

run: build
	./bin/go-blockchain

test:
	go test ./...