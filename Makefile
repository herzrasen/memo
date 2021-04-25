bin = memo

all: test

clean:
	rm -f $(bin)
	rm -f coverage.out

init:
	go mod download

test:
	go test -coverprofile=coverage.out -covermode=atomic -v ./...