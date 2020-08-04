FROM golang:1.15rc1-alpine3.12 as builder

RUN apk add --no-cache bash git

# Copy the code from the host and compile it
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download

ADD . ./
RUN [ "$(uname)" = Darwin ] && system=darwin || system=linux; \
    arch=amd64; \
    ./ci/go-build.sh --os ${system} --arch ${arch}

# final stage
FROM alpine:3.12
WORKDIR /app
COPY --from=builder /go/src/app/goapp /usr/local/bin/nats-seeder
ENTRYPOINT ["/usr/local/bin/nats-seeder"]
CMD ["mq"]
