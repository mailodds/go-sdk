FROM docker.io/library/golang:1.23-alpine
RUN apk add --no-cache curl
WORKDIR /sdk
COPY . .
RUN go build ./...
WORKDIR /smoke
RUN go mod init smoke && cp /sdk/ci/smoke_test.go main.go
RUN go build -o /smoke_test main.go
CMD ["/smoke_test"]
