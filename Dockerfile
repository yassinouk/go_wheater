# build stage
FROM golang:1.20-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main .
# final stage
FROM alpine:3.18 AS final
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]