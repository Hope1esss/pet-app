FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o pet-app ./cmd/main.go

CMD ["./pet-app"]