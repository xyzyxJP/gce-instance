FROM golang:1.19 AS builder

WORKDIR /work

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app .

FROM ubuntu:22.04

WORKDIR /work

RUN apt-get update && apt-get install -y --no-install-recommends \
        ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /work/app .

EXPOSE $PORT

ENTRYPOINT [ "/work/app" ]
