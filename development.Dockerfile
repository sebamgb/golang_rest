FROM golang:latest

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0 \
    GO111MODULE=on

RUN apt-get update

WORKDIR /go/src/app

COPY ["go.mod","go.sum", "./"]

RUN go mod download

COPY [".", "."]

CMD ["go", "run", "main.go"]