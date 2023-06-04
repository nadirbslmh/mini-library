FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o dist

EXPOSE 8083

ENTRYPOINT [ "./dist" ]