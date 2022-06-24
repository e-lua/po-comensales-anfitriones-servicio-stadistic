FROM golang:1.17.2

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN make build

EXPOSE 4320

CMD ["./ms-api"]
