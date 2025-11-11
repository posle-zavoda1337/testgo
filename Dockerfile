FROM golang:1.25.4-alpine3.22 as builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o myapp main.go

FROM alpine:3.18

RUN adduser -D -u 10001 appuser 

WORKDIR /app
COPY --from=builder /app/myapp .

USER appuser

EXPOSE 8080
ENTRYPOINT ["./myapp"]