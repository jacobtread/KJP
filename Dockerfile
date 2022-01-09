FROM golang:1.17.5-alpine3.15
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o app .
EXPOSE 4000
CMD ["/app/app"]