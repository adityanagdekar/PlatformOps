FROM golang:1.22-alpine AS build

WORKDIR /src

COPY go.mod ./
COPY main.go ./

RUN go build -o /out/sample-api .

FROM alpine:3.20

RUN adduser -D -H -u 10001 sampleapi

USER sampleapi
WORKDIR /app

COPY --from=build /out/sample-api /app/sample-api

EXPOSE 8080

ENTRYPOINT ["/app/sample-api"]
