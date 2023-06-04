FROM alpine:latest

WORKDIR /app

COPY bookApp /app

RUN apk add libc6-compat

EXPOSE 8083

ENTRYPOINT ["./bookApp"]
