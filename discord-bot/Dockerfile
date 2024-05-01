FROM golang:1.21.6

WORKDIR /app

COPY . .

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=1 GOOS=linux go build -buildvcs=false -o /discord-bot

EXPOSE 8080

CMD ["/discord-bot"]
