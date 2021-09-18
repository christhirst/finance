FROM golang:latest
ENV API_KEY_ID ${{ secrets.SECRET_KEY }}
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main ./cmd/web/
EXPOSE 8081
CMD ["/app/main"]