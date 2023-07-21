FROM golang:1.20-bookworm
WORKDIR /app
COPY goapp/go.mod .
COPY goapp/go.sum .
RUN go install github.com/cosmtrek/air@latest
RUN go get ./...
CMD go run .
