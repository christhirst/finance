FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main main.go
EXPOSE 8081
CMD ["/app/main"]
