FROM golang:1.17 AS builder

WORKDIR /go/src/app

COPY . .

# RUN go mod init && \
#     go build sum.go

CMD ["go test -coverprofile=coverage.out"]