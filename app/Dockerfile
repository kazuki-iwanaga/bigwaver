FROM golang:1.21-bullseye as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy