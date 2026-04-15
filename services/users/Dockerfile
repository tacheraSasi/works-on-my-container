FROM golang:1.22-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o users-service ./services/users
EXPOSE 8081
CMD ["./users-service"]