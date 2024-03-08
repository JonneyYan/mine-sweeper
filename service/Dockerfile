#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /usr/src/app

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN  go build -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder //usr/src/app/main .

EXPOSE 3001

CMD ["./main"]
