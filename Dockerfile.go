FROM golang:1.20-bookworm
WORKDIR /app
COPY goapp/go.mod .
COPY goapp/go.sum .
RUN go get ./...
CMD go run .
