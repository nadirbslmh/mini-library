AUTH_BINARY=authApp
BOOK_BINARY=bookApp
RENT_BINARY=rentApp
LIB_BINARY=libApp
LISTENER_BINARY=listenerApp
LOG_BINARY=logApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker compose up -d
	@echo "Docker images started!"

## up_build: stops docker compose (if running), builds all projects and starts docker compose
up_build: build_auth build_book build_rent build_lib build_listener build_log
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"

## build_auth: builds the auth binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd ./auth && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} cmd/*.go
	@echo "Done!"

## build_book: builds the book binary as a linux executable
build_book:
	@echo "Building book binary..."
	cd ./book && env GOOS=linux CGO_ENABLED=0 go build -o ${BOOK_BINARY} cmd/*.go
	@echo "Done!"

## build_rent: builds the rent binary as a linux executable
build_rent:
	@echo "Building rent binary..."
	cd ./rent && env GOOS=linux CGO_ENABLED=0 go build -o ${RENT_BINARY} cmd/*.go
	@echo "Done!"

## build_lib: builds the library binary as a linux executable
build_lib:
	@echo "Building library binary..."
	cd ./library && env GOOS=linux CC=/usr/bin/musl-gcc go build --ldflags '-linkmode external -extldflags "-static"' -tags musl -o ${LIB_BINARY} cmd/*.go
	@echo "Done!"

## build_listener: builds the listener binary as a linux executable
build_listener:
	@echo "Building listener binary..."
	cd ./listener && env GOOS=linux CC=/usr/bin/musl-gcc go build --ldflags '-linkmode external -extldflags "-static"' -tags musl -o ${LISTENER_BINARY} cmd/*.go
	@echo "Done!"

## build_log: builds the log binary as a linux executable
build_log:
	@echo "Building logging binary..."
	cd ./logging && env GOOS=linux CGO_ENABLED=0 go build -o ${LOG_BINARY} cmd/*.go
	@echo "Done!"