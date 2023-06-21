FROM golang:latest as build
LABEL name = devstack

WORKDIR /root

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
EXPOSE 6969


ENV DB_URL=localhost:5432/postgres:postgres@testdb
ENV DS_DB=postgres
ENV DB_NAME=testdb
ENV PORT = 8080


ENV  POSTGRES_USER=postgres
ENV  POSTGRES_PASSWORD=postgres
ENV  POSTGRES_URI=localhost
ENV  POSTGRES_PORT=5432

ENV PORT=6969

RUN go build  ./cmd/app/main.go

WORKDIR /app
COPY --from=build /root/main /app/

CMD ["./main"]

