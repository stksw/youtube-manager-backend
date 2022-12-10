FROM golang:1.19.2 as dev

RUN apt-get update

WORKDIR /app
RUN go install github.com/cosmtrek/air@latest

CMD ["air"]