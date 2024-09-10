FROM golang:1.23.1-bookworm

WORKDIR /app

COPY *.go /app

COPY . .

RUN go mod download

RUN cd cmd && go build -o /github.com/zachdehooge/radar_database

CMD [ "/github.com/zachdehooge/radar_database" ]
