FROM golang:1.18.3

WORKDIR /app

RUN apt-get update && \
    curl -sSL https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add - && \
    echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ bionic main" > /etc/apt/sources.list.d/migrate.list && \
    apt-get update && apt-get install -y migrate

COPY ./go.mod ./
COPY ./go.sum ./

COPY docker/start.sh ./

RUN go mod download

COPY . .

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go build -o ./out/groups ./cmd/main.go

EXPOSE 8081

RUN chmod +x /app/start.sh

CMD ["./start.sh"]