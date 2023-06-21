FROM golang:1.20-alpine
WORKDIR /app
COPY . .
EXPOSE 9090
RUN go build  ./cmd/app/main.go
ENTRYPOINT ["./main"]