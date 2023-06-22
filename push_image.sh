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

# logging

# listener