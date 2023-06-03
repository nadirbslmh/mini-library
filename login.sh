#!/bin/sh

token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjg1Nzc5OTU2fQ.8z2DzvDWuD7dzSD-A7-OpGwo1XzaWSOGpLtWToUwHdM"

# register
# curl -XPOST -H "Content-type: application/json" -d '{"email":"test@test.com","password":"blah"}' 'http://localhost:8080/api/v1/register'

# login
# curl -XPOST -H "Content-type: application/json" -d '{"email":"test@test.com","password":"blah"}' 'http://localhost:8080/api/v1/login'

# echo "test get books / others..."
# curl -XGET -H "Content-type: application/json" -H "Authorization: Bearer ${token}" 'http://localhost:8080/api/v1/books'