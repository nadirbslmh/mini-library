#!/bin/sh

token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjg1Nzg2OTE0fQ.JvBWGCU9R8nwz0E6R_FLcI9qXM5tLY9f7UonIu1fCoY"

# register
# curl -XPOST -H "Content-type: application/json" -d '{"email":"kzuna@test.com","password":"blah"}' 'http://localhost:8080/api/v1/register'

# login
# curl -XPOST -H "Content-type: application/json" -d '{"email":"test@test.com","password":"blah"}' 'http://localhost:8080/api/v1/login'

echo "test create rent"
curl -XGET -H "Content-type: application/json" -H "Authorization: Bearer ${token}" -d '{"book_id":3}' 'http://localhost:8080/api/v1/rents'