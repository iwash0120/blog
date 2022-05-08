FROM golang:1.16.5 as release_builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM alpine:latest as release
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=release_builder /app .
CMD ["./app"]

FROM golang:1.16.5 as debug
RUN go get github.com/cosmtrek/air@v1.29.0
WORKDIR /app
CMD ["air", "-c", ".air.toml"]
