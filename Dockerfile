FROM golang:1.21.6 AS builder

WORKDIR /app

COPY . .

#RUN go mod tidy

#RUN go build -o memoo-backend

EXPOSE 8080

CMD ["./memoo-backend"]
