FROM alpine:latest

RUN apk update && apk add --no-cache bash ca-certificates git gcc g++ musl-dev librdkafka-dev pkgconf

WORKDIR /app

COPY libApp /app

EXPOSE 8080

ENTRYPOINT ["./libApp"]