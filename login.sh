#!/bin/sh

token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjg1NzgxNDI3fQ.XkoAR6D3Pzcx3ezU1B7JHGy-i5vhv1mDH3P8E9xZe6M"

# register
# curl -XPOST -H "Content-type: application/json" -d '{"email":"nezuko@test.com","password":"blah"}' 'http://localhost:8080/api/v1/register'

# login
# curl -XPOST -H "Content-type: application/json" -d '{"email":"test@test.com","password":"blah"}' 'http://localhost:8080/api/v1/login'

# echo "test create rent"
# curl -XPOST -H "Content-type: application/json" -H "Authorization: Bearer ${token}" -d '{"book_id":11}' 'http://localhost:8080/api/v1/rents'