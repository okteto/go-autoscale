FROM golang:buster as builder

WORKDIR /app
ADD . .
RUN go build -o /usr/local/bin/calculate

EXPOSE 8080
CMD ["/usr/local/bin/calculate"]
