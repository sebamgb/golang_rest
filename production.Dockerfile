ARG GO_VERSION=1.19.1

FROM golang:${GO_VERSION}-alpine3.16 as builder

ENV GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on

RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates && update-ca-certificates

WORKDIR /go/src

COPY ["./go.mod", "./go.sum", "./"]

RUN go mod download

COPY [".", "."]

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /go/bin/rest-go-ws

# production

FROM scratch as runner

COPY --from=builder ["/etc/ssl/certs/ca-certificates.crt","/etc/ssl/certs"]

COPY ["./.env", "./"]

COPY --from=builder ["/go/bin/rest-go-ws", "/go/bin/rest-go-ws"]

EXPOSE 5050

ENTRYPOINT ["/go/bin/rest-go-ws"]