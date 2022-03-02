FROM golang:1.17-alpine

RUN  export GO111MODULE=auto

WORKDIR /app

COPY  src/go.mod .
COPY  src/go.sum .

RUN go mod download && go mod verify

COPY src/ ./

RUN go build cmd/main.go

EXPOSE 8080

CMD ["./main"]
