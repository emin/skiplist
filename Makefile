build:
	go build .

test:
	go test -v ./...

cover:
	go test -race -coverprofile=coverage.txt -covermode=atomic


show-cover: cover
	rm -f coverage.txt