FROM golang:1.18

WORKDIR /app
COPY main.go .
COPY go.mod .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /server main.go

EXPOSE 8080

CMD [ "/server" ]