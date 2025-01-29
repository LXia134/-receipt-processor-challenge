FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .

FROM golang:1.23
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]