FROM golang:1.18.4

EXPOSE 8000

WORKDIR /app

COPY ./src /app

RUN go get github.com/go-sql-driver/mysql

CMD ["go","run","main.go"]
