# develop
FROM golang:1.20.13-alpine AS develop

ENV CGO_ENABLED 0

WORKDIR /app

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod tidy


# builder
FROM golang:1.20.13-alpine as builder

WORKDIR /app

RUN apk update && apk add git
COPY go.mod go.sum main.go ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin


# production
FROM scratch as production

COPY --from=builder /app/bin /app/bin

EXPOSE 8080
CMD ["/app/bin"]