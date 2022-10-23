FROM golang:1.18

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN make build

EXPOSE 4320

CMD ["./ms-api"]
