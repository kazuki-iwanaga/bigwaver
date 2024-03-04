FROM golang:1.21-bullseye as builder

WORKDIR /app

COPY go.mod go.sum main.go publisher.go validate.go ./

RUN go mod tidy \
    && CGO_ENABLED=0 GOOS=linux go build -o /app/app

FROM scratch as production

WORKDIR /app

COPY --from=builder /app/app /app/app

EXPOSE 8080

CMD ["/app/app"]