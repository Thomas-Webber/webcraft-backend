BINARY_NAME=webcraft-backend

build:
	go build -o /tmp/$(BINARY_NAME) -v ./main.go

run: build
	/tmp/$(BINARY_NAME)

clean:
	rm -f /tmp/$(BINARY_NAME)

lint: # goreportcard.com, github.com/golang/go/wiki/CodeReviewComments
	go fmt ./...
	golint ./...
