FROM alpine:latest

WORKDIR /app

COPY authApp /app

RUN apk add libc6-compat

EXPOSE 8083

ENTRYPOINT ["./authApp"]
