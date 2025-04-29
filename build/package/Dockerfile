FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata curl

WORKDIR /build

ADD go.mod .

ADD go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /app/geo-forecast ./cmd/geo-forecast/geo-forecast.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

ENV TZ Etc/UTC

WORKDIR /app

COPY --from=builder /app/geo-forecast /app/geo-forecast

EXPOSE 8080

CMD ["./geo-forecast"]

