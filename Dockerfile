FROM golang:1.11 as goimage
ENV SRC=/go/src/
RUN mkdir -p /go/src/
WORKDIR /go/src/go_docker
RUN git clone https://github.com/simenstoa/trening-backend.git /go/src/go_docker/ \
  && CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go get ./
RUN go build -o bin/go_docker

ENTRYPOINT bin/go_docker
EXPOSE 8080