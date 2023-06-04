FROM alpine:latest

WORKDIR /app

COPY rentApp /app

RUN apk add libc6-compat

EXPOSE 8082

ENTRYPOINT ["./rentApp"]
