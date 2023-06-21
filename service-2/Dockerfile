FROM golang:1.20-alpine
WORKDIR /app
COPY . .
EXPOSE 6969
RUN go build  ./cmd/app/main.go
ENTRYPOINT ["./main"]