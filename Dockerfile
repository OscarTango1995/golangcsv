FROM golang:1.15.5-buster
LABEL maintainer="Ovais Tariq <ovais.tariq@hotmail.com>"


RUN apt-get update && apt-get install -y default-mysql-client
RUN go get github.com/go-sql-driver/mysql

RUN mkdir -p /app/task
ADD . /app/task/
WORKDIR /app/task


RUN go build -o main .
CMD ["/app/task/main"]