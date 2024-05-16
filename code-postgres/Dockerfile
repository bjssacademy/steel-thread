FROM golang:1.22.3-alpine3.19

RUN mkdir /app

ADD . /app

WORKDIR /app

ENV GOOSE_DRIVER=postgres
ENV GOOSE_DBSTRING=user=postgres password=YOURPASSWORD dbname=acme sslmode=disable host=host.docker.internal

EXPOSE 8080

RUN go build -o main .

CMD ["/app/main"]