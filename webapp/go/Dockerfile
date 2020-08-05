FROM golang:1.12

RUN apt-get update && apt-get -y install mariadb-client

ENV GO111MODULE=on

WORKDIR /go/src/webapp
RUN go build -o main *.go
RUN chmod +x main
CMD ["main"]
