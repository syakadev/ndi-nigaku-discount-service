FROM golang:1.23.0-alpine AS builder

RUN apk update && apk add --no-cache \
    git build-base zeromq-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

ENV CGO_ENABLED=1
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -o binary .

FROM alpine:latest

RUN apk add --no-cache zeromq ca-certificates

WORKDIR /app
COPY --from=builder /app/binary /app/binary

RUN adduser -D ndi
USER ndi

ENTRYPOINT ["/app/binary"]