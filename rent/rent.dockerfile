FROM alpine:latest

WORKDIR /app

COPY rentApp /app

EXPOSE 8082

ENTRYPOINT ["./rentApp"]
