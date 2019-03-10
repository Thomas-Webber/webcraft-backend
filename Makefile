BINARY_NAME=webcraft-backend

build:
	go build -o /tmp/$(BINARY_NAME) -v ./main.go

run: build
	/tmp/$(BINARY_NAME)

clean:
	rm -f /tmp/$(BINARY_NAME)

dep:
	go get -d -v

lint: # goreportcard.com, github.com/golang/go/wiki/CodeReviewComments
	go fmt ./...
	golint ./...
