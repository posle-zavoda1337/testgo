FROM golang:1.25.4-alpine3.22 as builder

WORKDIR /app

COPY . .

RUN go build -o myapp

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/myapp .

ENTRYPOINT ["./myapp"]