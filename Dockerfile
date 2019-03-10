# https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
#
# Produce a docker image for production

FROM golang as builder
COPY . $GOPATH/src/webcraft
WORKDIR $GOPATH/src/webcraft
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/webcraft

FROM scratch
COPY --from=builder /go/bin/webcraft /go/bin/webcraft
ENTRYPOINT ["/go/bin/webcraft"]