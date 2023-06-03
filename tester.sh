#!/bin/sh

# get all books
# curl -XGET 'http://localhost:8080/api/v1/books'

# # create book
# curl -XPOST -H "Content-type: application/json" -d '{"title":"other book","description":"blah","author":"jean eric pernandez"}' 'http://localhost:8080/api/v1/books'

# # get book by id
curl -XGET 'http://localhost:8080/api/v1/books/1'

# get all rents
# curl -XGET 'http://localhost:8080/api/v1/rents'

# create rent
# curl -XPOST -H "Content-type: application/json" -d '{"user_id":1,"book_id":1,"book_title":"belum dinamis"}' 'http://localhost:8080/api/v1/rents'