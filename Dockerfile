FROM golang:1.9

MAINTAINER P.Mauran (maurakiri@gmail.com)

# install dependencies
RUN go get github.com/codegangsta/negroni
RUN	go get github.com/dgrijalva/jwt-go
RUN go get github.com/dgrijalva/jwt-go/request
RUN go get github.com/gorilla/mux

# env
ENV HPORT 8002
ENV ID_RSA .keys/app.rsa
ENV ID_RSA_PUB .keys/app.rsa.pub


# copy app
ADD . /app
WORKDIR /app

# build
RUN go build -o build/gojwt src/*.go

# server running port
EXPOSE 8002

# .keys volume
VOLUME ["/app/.keys"]

ENTRYPOINT ["/app/docker-entrypoint.sh"]
