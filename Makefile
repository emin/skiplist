build:
	go build .

test:
	go test -v ./...

cover:
	go test -race -coverprofile=coverage.txt -covermode=atomic

bench:
	go test -bench=.

show-cover:
	go test -cover