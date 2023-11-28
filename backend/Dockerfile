FROM golang:1.21-alpine3.17

WORKDIR /usr/src/backend

COPY . .

RUN make

EXPOSE 443

CMD ["./main"]

