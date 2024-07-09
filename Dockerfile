ARG GO_VERSION=1.21.0
FROM golang:${GO_VERSION}-bookworm as builder

RUN mkdir /Users/aa/bucket
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
ENV DATABASE_URL=gnr
RUN go build -v -o /usr/src/app/run-app .

FROM debian:bookworm

COPY --from=builder /usr/src/app/run-app /usr/local/bin/
CMD ["/usr/local/bin/run-app", "run", "8080"]
