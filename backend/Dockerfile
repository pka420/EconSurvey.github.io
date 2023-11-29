FROM golang:1.21-alpine3.17

WORKDIR /go/delivery

COPY . /go/delivery/

RUN apk add --no-cache make build-base sqlite
RUN make

EXPOSE 8080

CMD ["./main"]

