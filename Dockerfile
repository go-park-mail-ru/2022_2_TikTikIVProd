FROM golang:1.16-alpine

WORKDIR /

COPY app ./app
COPY app ./app

WORKDIR /app/

RUN go mod download

RUN go build -o /go_server

EXPOSE 8080

CMD [ "/go_server" ]