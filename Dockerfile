ARG go="golang:1.22.6-alpine3.20"
ARG base="alpine:3.20"

FROM --platform=$BUILDPLATFORM ${go} AS builder

RUN apk add --no-cache bash git

ARG TARGETPLATFORM
ARG BUILDPLATFORM

# Copy the code from the host and compile it
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download

ADD . ./

RUN [ "$(uname)" = Darwin ] && system=darwin || system=linux; \
    ./ci/go-build.sh --os ${system} --arch $(echo $TARGETPLATFORM  | cut -d/ -f2)

# install nats cli
RUN go install github.com/nats-io/natscli/nats@latest

# final stage
FROM ${base}
WORKDIR /app
COPY --from=builder /go/src/app/goapp /usr/local/bin/nats-seeder
COPY --from=builder /go/bin/nats /usr/local/bin/nats
ENTRYPOINT ["/usr/local/bin/nats-seeder"]
CMD ["help"]
