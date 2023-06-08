# mini-library

A simple book library application with microservices architecture. Written in Go.

## Additional Notes

This repository uses gRPC to communicate with other services.

## How to Use

1. Clone this repository.

2. Make sure `make` and `docker` are installed.

3. Build the application.

```sh
make up_build
```

4. The endpoints are available in this base url: `http://localhost:8080/api/v1`.

5. Stop the application.

```sh
make down
```
