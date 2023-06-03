#!/bin/sh

token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywiZXhwIjoxNjg1Nzg1OTg4fQ.59PXeSq1Dvk8MBPPoOrZYpZQBBn_zEo9Ki38H2F8qv0"

# register
# curl -XPOST -H "Content-type: application/json" -d '{"email":"nezuko@test.com","password":"blah"}' 'http://localhost:8080/api/v1/register'

# login
# curl -XPOST -H "Content-type: application/json" -d '{"email":"nezuko@test.com","password":"blah"}' 'http://localhost:8080/api/v1/login'

# echo "test create rent"
curl -XGET -H "Content-type: application/json" -H "Authorization: Bearer ${token}" -d '{"book_id":11}' 'http://localhost:8080/api/v1/rents'