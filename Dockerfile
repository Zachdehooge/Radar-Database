FROM golang:1.24.4-bookworm

WORKDIR /app

COPY *.go /app

COPY . .

RUN go mod download 

RUN go build -o /github.com/zachdehooge/radar-database

CMD [ "/github.com/zachdehooge/radar-database" ]
