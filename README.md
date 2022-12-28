# Go-gRPC

Just simple client server app using go and grpc

## Requirements

1. Go
2. Protobuf
3. Go-protoc extension

## What do I learn

1. gRPC is schema on write so we must defined schema first then generate the code
2. Protobuf is the serialization and gRPC is the protocol
   1. use `protoc --go_out` to generate the model
   2. use `prococ --go-grpc_out` to generate the method
3. gRPC use params numbering means the parameters must have matching number start from one
4. THe service in gRPC **cannot** use primitive type, must `wrapped` it to another type such as `GarageUserId` and etc
5. The service in gRPC must have parameter and reuturn value, use `google.protobuf.empty` to handle such case
6. Genereated service interface must wrapped with another struct and then implement the method

## References

1. [Dasar Pemrogramman golang novalagung](https://dasarpemrogramangolang.novalagung.com/C-golang-grpc-protobuf.html)
2. [Issue implemented](https://github.com/grpc/grpc-go/issues/3794)
