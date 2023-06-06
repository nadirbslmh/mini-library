# generate code for gorm additional types
# protoc -I=api --go_out=./pkg --go-grpc_out=./pkg gorm.proto

# generate code for proto files with imports [use this!]
# protoc -I=api --go_out=./pkg --go-grpc_out=./pkg gorm.proto auth.proto book.proto


# generate code for auth grpc
# protoc -I=api --go_out=./pkg --go-grpc_out=./pkg auth.proto

# generate code for book grpc
# protoc -I=api --go_out=./pkg --go-grpc_out=./pkg book.proto

# generate code for rent grpc
# protoc -I=api --go_out=./pkg --go-grpc_out=./pkg rent.proto