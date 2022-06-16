FROM golang:alpine3.16

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-steam-item-tracker

CMD [ "/docker-steam-item-tracker" ]
