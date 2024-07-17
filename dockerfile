FROM golang:alpine

RUN mkdir app

WORKDIR /app

ADD . /app

RUN go build -o server .

CMD [ "./server" ]