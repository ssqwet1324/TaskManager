FROM golang:1.24.5-alpine

WORKDIR /app

COPY . .

RUN go build -o application main.go

EXPOSE 8081

CMD ["./application"]