FROM golang:latest as build
LABEL name = devstack

WORKDIR /root

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
EXPOSE 6969


ENV DS_DB=postgres
ENV DB_NAME=testdb

ENV  DB_USER=postgres
ENV  DB_PASSWORD=postgres
ENV  DB_URL=localhost
ENV  DB_PORT=5432

ENV PORT=6969

RUN go build  ./cmd/app/main.go

WORKDIR /app
COPY --from=build /root/main /app/

CMD ["./main"]

