# library service
docker build -t library-service:v1 -f library.dockerfile .
docker tag library-service:v1 nadirbasalamah/library-service:v1
docker push nadirbasalamah/library-service:v1

# auth

# book

# rent

# logging

# listener