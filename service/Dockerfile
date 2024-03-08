#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /usr/src/app

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go app
RUN  go build -o main .

#final stage
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder //usr/src/app/main .

# Expose port 8888 to the outside world
EXPOSE 3030

#Command to run the executable
CMD ["./main"]
