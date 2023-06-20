# mini-library

A simple book library application with microservices architecture. Written in Go.

## How to Use

1. Clone this repository.

2. Make sure `make` and `docker` are installed.

3. There is a library that requires C language feature in this repository (example: `confluent-kafka-go`). Make sure [`musl`](https://musl.libc.org/) is installed.

4. Copy the configuration file. Then fill the required configurations.

```sh
cp .env.example ./library/.env
```

5. Build the application.

```sh
make up_build
```

6. The endpoints are available in this base url: `http://localhost:1323/api/v1`.

7. Stop the application.

```sh
make down
```

## Perform Static Analysis

1. Start the SonarQube with Docker.

```sh
docker run -d --name sonarqube -e SONAR_ES_BOOTSTRAP_CHECKS_DISABLE=true -m 2g -p 9000:9000 sonarqube:9.9.1-community
```

2. Run the test.

- Make sure to replace `YOUR_TOKEN` with the generated token from the SonarQube.
- Make sure to replace `/path/to/app` with the location of your project.

```sh
docker run \
--rm \
-e SONAR_HOST_URL="http://localhost:9000" \
-e SONAR_LOGIN="YOUR_TOKEN" \
-v "/path/to/app:/usr/src" \
--network host \
-m 1g \
sonarsource/sonar-scanner-cli
```

Or use the provided command in each service.

| **Command**         | **Description**                                |
| ------------------- | ---------------------------------------------- |
| `./auth/analyze.sh` | Perform static code analysis for auth service. |
| `./book/analyze.sh` | Perform static code analysis for book service. |
| `./rent/analyze.sh` | Perform static code analysis for rent service. |
