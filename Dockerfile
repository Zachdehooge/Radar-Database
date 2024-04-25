FROM golang:1.22.2-bookworm

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /github.com/zachdehooge/radar_database

CMD [ "/github.com/zachdehooge/radar_database" ]
