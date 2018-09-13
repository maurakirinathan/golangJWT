FROM golang:1.9

MAINTAINER Eranga Bandara (erangaeb@gmail.com)

# install dependencies
RUN go get github.com/gocql/gocql
RUN	go get github.com/gorilla/mux
RUN go get github.com/Shopify/sarama
RUN go get github.com/wvanbergen/kafka/consumergroup
RUN go get github.com/gorilla/handlers

# env
#ENV HPORT 2181

# copy app
ADD . /app
WORKDIR /app

# build
RUN go build -o build/senz src/*.go

# server running port
EXPOSE 8002

# .keys volume
VOLUME ["/app/.keys"]

ENTRYPOINT ["/app/docker-entrypoint.sh"]
