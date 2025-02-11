FROM bufbuild/buf:latest as bufbuild
WORKDIR /app

COPY . .

RUN buf generate \
        --path api/schema/probe/healt_check.proto

# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.22 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

COPY --from=bufbuild /app/. /app/.

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o /app/app /app/cmd/*.go

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/* && mkdir /app


# Set working directory
WORKDIR /app

# Copy the binary from the "builder" stage
COPY --from=builder /app/app /app/app
# Copy config directory for the file

# Expose port
EXPOSE 8085

CMD ["/app/app"]