FROM alpine:latest

WORKDIR /app

COPY authApp /app

EXPOSE 8083

ENTRYPOINT ["./authApp"]
