FROM golang:latest as builder

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0 \
    GO111MODULE=on

RUN apt-get update

WORKDIR /go/src/app

COPY ["go.mod", "."]

RUN go mod download

COPY [".", "."]

RUN go build -v main.go

# production

FROM scratch

COPY --from=builder ["/go/src/app","."]

ENTRYPOINT ["./main"]