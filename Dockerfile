FROM golang:1.19-alpine AS builder

LABEL maintainer="Dhruv Vavliya <dhruvvavliya79@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o main server.go

EXPOSE 80

CMD ["./main"]