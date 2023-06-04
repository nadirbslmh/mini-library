FROM alpine:latest

WORKDIR /app

COPY libApp /app

EXPOSE 8080

ENTRYPOINT ["./libApp"]
