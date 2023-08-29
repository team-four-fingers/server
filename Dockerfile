# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
FROM golang:latest as builder

# Create and change to the app directory.
WORKDIR /server
COPY . .

RUN go mod download

# Build the binary.
WORKDIR /server/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -ldflags="-w -s" -o /go/bin/server
#RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -mod=readonly -v -o server /cmd/main.go

FROM debian:stable-20230320-slim

RUN apt-get update \
    && apt-get install -y ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /go/bin /app

EXPOSE 8080

CMD ["/app/server"]
