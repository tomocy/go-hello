FROM golang:latest as build

WORKDIR /go/src/github.com/tomocy/go-hello
COPY . .

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0
RUN go build -o app .

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=build /go/src/github.com/tomocy/go-hello/app /app

ENTRYPOINT ["/app"]