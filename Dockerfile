FROM golang:1.11 AS builder
RUN apt-get install -y ca-certificates
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/gotzapi

FROM alpine:latest
RUN apk --no-cache add ca-certificates libc6-compat
COPY --from=builder /go/bin/gotzapi /go/bin/gotzapi
ENTRYPOINT ["/go/bin/gotzapi"]