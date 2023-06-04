FROM alpine:latest

WORKDIR /app

COPY libApp /app

RUN apk add libc6-compat

EXPOSE 8083

ENTRYPOINT ["./libApp"]
