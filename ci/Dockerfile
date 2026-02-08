FROM docker.io/library/golang:1.23-alpine
WORKDIR /sdk
COPY . .
RUN go build ./...
WORKDIR /smoke
COPY ci/smoke_test.go main.go
RUN go mod init smoke && go mod edit -require github.com/mailodds/go-sdk@v0.0.0 -replace github.com/mailodds/go-sdk=/sdk && go mod tidy && go build -o /smoke_test main.go
CMD ["/smoke_test"]
