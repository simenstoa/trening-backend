FROM golang:1.11 as goimage
RUN mkdir -p /go/src/
WORKDIR /go/src/go_docker
COPY . .
RUN go get ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/go_docker

FROM alpine:3.8 as baseimagealp
RUN apk add --no-cache bash
WORKDIR /docker/bin
COPY --from=goimage /go/src/go_docker/bin/ ./
ENTRYPOINT /docker/bin/go_docker
EXPOSE 8080
