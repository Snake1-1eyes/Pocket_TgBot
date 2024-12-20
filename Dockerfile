FROM golang:1.23.4-alpine3.21 AS builder

COPY . /Pocket_TgBot/
WORKDIR /Pocket_TgBot/

RUN go mod download
RUN GOOS=linux go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /Pocket_TgBot/bin/bot .
COPY --from=builder /Pocket_TgBot/configs configs/

EXPOSE 80

CMD ["./bot"]
