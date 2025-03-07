FROM golang:1.24-alpine
WORKDIR /simple-http-server
COPY . .
RUN go build -o server
EXPOSE 8080
CMD ["/simple-http-server/server"]