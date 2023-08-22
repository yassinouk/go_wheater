# build stage
FROM golang:1.20-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main .
# final stage using scratch instead of alpine to reduce the weight of the image
FROM scratch AS final
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]