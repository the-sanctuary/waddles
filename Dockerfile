# compile waddles inside a container
FROM golang:1.16-alpine AS builder

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

# create a barebones container to actually run waddles in
FROM scratch

# copy our static binary
COPY --from=builder /build/bin/waddles /

# copy ca-certificat
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


# start the bot
ENTRYPOINT ["/waddles"]
