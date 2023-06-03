#!/bin/sh

token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjg1Nzc5OTU2fQ.8z2DzvDWuD7dzSD-A7-OpGwo1XzaWSOGpLtWToUwHdM"

# get all books
# curl -XGET -H "Authorization: Bearer ${token}" 'http://localhost:8080/api/v1/books'

# # create book
# curl -XPOST -H "Content-type: application/json" -H "Authorization: Bearer ${token}" -d '{"title":"demon slayer","description":"blah","author":"jean eric pernandez"}' 'http://localhost:8080/api/v1/books'

# # get book by id
# curl -XGET -H "Authorization: Bearer ${token}" 'http://localhost:8080/api/v1/books/11'

# get all rents
# curl -XGET -H "Authorization: Bearer ${token}" 'http://localhost:8080/api/v1/rents'

# create rent
# curl -XPOST -H "Content-type: application/json" -H "Authorization: Bearer ${token}" -d '{"user_id":1,"book_id":11}' 'http://localhost:8080/api/v1/rents'