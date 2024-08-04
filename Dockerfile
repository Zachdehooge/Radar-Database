FROM golang:1.23rc2-bookworm

WORKDIR /app

COPY *.go /app

COPY . .

RUN go mod download

RUN go build -o /github.com/zachdehooge/radar_database

CMD [ "/github.com/zachdehooge/radar_database" ]
