FROM golang:1.19

WORKDIR /

COPY app ./app
COPY images ./images

WORKDIR /app

RUN go mod download
RUN go mod tidy
RUN go build cmd/main.go

WORKDIR /

EXPOSE 8080

CMD [ "./app/main" ]