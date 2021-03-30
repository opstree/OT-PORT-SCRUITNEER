FROM golang:1.16-alpine

RUN apk add --no-cache git nmap

# Set the Current Working Directory inside the container
WORKDIR /app/port-scanner

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o port-scanner

# Run the binary program produced by `go install`
CMD ["/app/port-scanner/port-scanner"]