# builder
FROM golang:1.20.13-alpine as builder

RUN apk update && apk add git

WORKDIR /tmp
RUN wget https://github.com/adnanh/webhook/releases/download/2.8.1/webhook-linux-arm64.tar.gz \
    && tar -xvf webhook-linux-arm64.tar.gz \
    && mv webhook-linux-arm64/webhook /usr/local/bin/webhook \
    && rm -rf webhook-linux-arm64.tar.gz webhook-linux-arm64

WORKDIR /app
COPY go.mod go.sum main.go ./
RUN go mod tidy \
    && CGO_ENABLED=0 GOOS=linux go build -o /app/publish

# production
# FROM scratch as production
FROM bash:5.1.16-alpine3.19 as production

WORKDIR /app

COPY hooks.yaml hooks.sh /app/
COPY --from=builder /usr/local/bin/webhook /usr/local/bin/webhook
COPY --from=builder /app/publish /app/publish

EXPOSE 8080
CMD ["webhook", "-hooks", "/app/hooks.yaml", "-port", "8080", "-verbose"]