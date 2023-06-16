# mini-library

A simple book library application with microservices architecture. Written in Go.

## How to Use

1. Clone this repository.

2. Make sure `make` and `docker` are installed.

3. There is a library that requires C language feature in this repository (example: `confluent-kafka-go`). Make sure [`musl`](https://musl.libc.org/) is installed.

4. Build the application.

```sh
make up_build
```

5. The endpoints are available in this base url: `http://localhost:1323/api/v1`.

6. Stop the application.

```sh
make down
```
