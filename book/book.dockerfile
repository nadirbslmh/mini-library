FROM alpine:latest

WORKDIR /app

COPY bookApp /app

EXPOSE 8081

ENTRYPOINT ["./bookApp"]
