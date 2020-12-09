# compile waddles inside a container
FROM golang:1.15-alpine AS builder

# # we need ca-certificates in our final container to make discord api requests
# RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /build

# download dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy source code to build dir
COPY . .

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# build our static binary
RUN go build -o /build/bin/waddles ./cmd/waddles/

# create a barebones container to actually run in
FROM scratch
# FROM golang:1.15-alpine

# copy our static binary
COPY --from=builder /build/bin/waddles /
# copy the config file
COPY waddles.toml ./
# copy ca-certificat
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


# start the bot
ENTRYPOINT ["/waddles"]