# library service
docker build -t library-service:v1 -f library.dockerfile .
docker tag library-service:v1 nadirbasalamah/library-service:v1
docker push nadirbasalamah/library-service:v1

# auth
docker build -t auth-service:v2 -f auth.dockerfile .
docker tag auth-service:v2 nadirbasalamah/auth-service:v2
docker push nadirbasalamah/auth-service:v2

# book
docker build -t book-service:v1 -f book.dockerfile .
docker tag book-service:v1 nadirbasalamah/book-service:v1
docker push nadirbasalamah/book-service:v1

# rent
docker build -t rent-service:v1 -f rent.dockerfile .
docker tag rent-service:v1 nadirbasalamah/rent-service:v1
docker push nadirbasalamah/rent-service:v1

# logging
docker build -t logging-service:v1 -f logging.dockerfile .
docker tag logging-service:v1 nadirbasalamah/logging-service:v1
docker push nadirbasalamah/logging-service:v1

# listener
docker build -t listener-service:v1 -f listener.dockerfile .
docker tag listener-service:v1 nadirbasalamah/listener-service:v1
docker push nadirbasalamah/listener-service:v1