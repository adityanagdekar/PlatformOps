FROM golang:1.22-alpine AS build

WORKDIR /src

COPY go.mod ./
COPY cmd ./cmd

RUN go build -o /out/platformops ./cmd/server

FROM alpine:3.20

RUN adduser -D -H -u 10001 platformops

USER platformops
WORKDIR /app

COPY --from=build /out/platformops /app/platformops

EXPOSE 8080

ENTRYPOINT ["/app/platformops"]

