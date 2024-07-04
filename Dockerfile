FROM golang:1.22.5-bookworm

WORKDIR /app

COPY *.go /app

COPY . .

RUN go mod download

RUN go build -o /github.com/zachdehooge/radar_database

CMD [ "/github.com/zachdehooge/radar_database" ]
