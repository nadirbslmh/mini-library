#!/bin/sh

# register
# curl -XPOST -H "Content-type: application/json" -d '{"email":"test@test.com","password":"blah"}' 'http://localhost:8080/api/v1/register'

# login
curl -XPOST -H "Content-type: application/json" -d '{"email":"test@test.com","password":"blah"}' 'http://localhost:8080/api/v1/login'

