FROM golang as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o webcraft

FROM scratch

ENV ENV=prod
ENV PORT=80

COPY --from=builder /app/webcraft webcraft
ENTRYPOINT ["/go/bin/webcraft"]