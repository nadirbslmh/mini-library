FROM alpine:latest

WORKDIR /app

COPY logApp /app

EXPOSE 8085

ENTRYPOINT ["./logApp"]
