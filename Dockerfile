FROM golang:1.19-alpine AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app .

FROM alpine:3.17.0

ENV GIN_MODE=release

RUN addgroup -S appuser && adduser -S appuser -G appuser
COPY --chown=appuser:appuser --from=builder /usr/src/app/app ./app

EXPOSE 8080
USER appuser

ENTRYPOINT [ "./app" ]
