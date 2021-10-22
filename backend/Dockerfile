# syntax=docker/dockerfile:1


FROM golang:1.16-alpine


WORKDIR /app

COPY go.mod ./
COPY go.sum ./


RUN go mod download

COPY src/*.go ./

RUN go build -o /go-chat


EXPOSE 8001

CMD [ "/go-chat" ]



