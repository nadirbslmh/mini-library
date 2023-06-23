# library service
docker build -t library-service:v1 -f library.dockerfile .
docker tag library-service:v1 dockerhub_username/library-service:v1
docker push dockerhub_username/library-service:v1

# auth
docker build -t auth-service:v2 -f auth.dockerfile .
docker tag auth-service:v2 dockerhub_username/auth-service:v2
docker push dockerhub_username/auth-service:v2

# book
docker build -t book-service:v1 -f book.dockerfile .
docker tag book-service:v1 dockerhub_username/book-service:v1
docker push dockerhub_username/book-service:v1

# rent
docker build -t rent-service:v1 -f rent.dockerfile .
docker tag rent-service:v1 dockerhub_username/rent-service:v1
docker push dockerhub_username/rent-service:v1

# logging
docker build -t logging-service:v1 -f logging.dockerfile .
docker tag logging-service:v1 dockerhub_username/logging-service:v1
docker push dockerhub_username/logging-service:v1

# listener
docker build -t listener-service:v1 -f listener.dockerfile .
docker tag listener-service:v1 dockerhub_username/listener-service:v1
docker push dockerhub_username/listener-service:v1

# cmd references
minikube addons enable ingress