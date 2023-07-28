FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go build -o domain /app/cmd/main.go

FROM alpine:latest
COPY --from=builder /app/rover /app/rover
COPY --from=builder /app/input.txt /app/input.txt
CMD ["/app/rover"]