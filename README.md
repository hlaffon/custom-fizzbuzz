# Custom-FizzBuzz

## Description

This application will return all numbers from 1 to 200 but for multiples of 3 it will print "Fizz" instead of the 
number and for the multiples of 5 it will print "Buzz".
Moreover, for numbers which are multiples of both 3 and 5 it will print "FizzBuzz".

In addition to this default behavior, the user can customize the app settings by defining:

- the maximum numbers that will be returned
- the first and second multiple to be replaced
- the aliases that will replace the two multiples

NB: If the user doesn't define one or more of the parameters, then the default value will be used.

## Pre-requisites

To run the application locally you need either :
- A local go installation (see https://go.dev/doc/install)
- Docker compose installed (see https://docs.docker.com/compose/install/)

## Usage

The application can be launched locally by running this command:
`go run cmd/main.go`

Or you can run it in a docker container:
`docker compose up -d
`

(To stop the container:
`docker compose down`)

We offer two ways of retrieving the requested data, depending on your needs and ability to implement a given client.
The first one is a http REST API with JSON format which is widely used, human-readable and easy to request
using tools like postman for instance. However, if you need high performance, quick data sending of large payloads,
or lighter data transfer to reduce your environmental impact, you can use the gRPC API with protobuf format. 

The http server will be running on port 4000 and the gRPC one on port 5000.


### Http

We provide a http server to request the REST API with the following endpoint:

`
http://localhost:4000/numbers/print
`

It can be called either with a GET method using query params or using a POST with a request body.

To retrieve the numbers you can use curl (or any http client provider such as postman).

#### GET

Using the GET method, you just need to specify the requested values with query parameters, for instance :

`
curl -v -X GET "http://localhost:4000/numbers/print?maxNumber=50&firstMultiple=3&secondMultiple=5&firstAlias=kiss&secondAlias=cool"
`


#### POST

If you prefer to specify a request body containing all the custom settings, you can use this json structure with a POST method:

```
{
MaxNumber      *int    `json:"maxNumber"`
FirstMultiple  *int    `json:"firstMultiple"`
SecondMultiple *int    `json:"secondMultiple"`
FirstAlias     *string `json:"firstAlias"`
SecondAlias    *string `json:"secondAlias"`
}
```

An example of request:

`
curl -v -X POST "http://localhost:4000/numbers/print" -H "Content-Type: application/json" -d '{"maxNumber":50, "firstMultiple":3, "secondMultiple":5, "firstAlias":"kiss", "secondAlias": "cool"}'
`


### gRPC

#### Protobuf generation
To regenerate the pb.go files from the proto message definition you need to have installed a protobuf compiler as well as the protoc-gen-go library:

`
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
`

`
sudo apt  install protobuf-compiler (for ubuntu)
`

Then you can run this command from the root of the project:

`
protoc --proto_path=proto proto/api.proto --go_out=pkg/pb --go-grpc_out=pkg/pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
`


#### How to request the API

Here is the protobuf messages definition :

```
message Request {
    google.protobuf.Int32Value MaxNumber = 1;
    google.protobuf.Int32Value FirstMultiple = 2;
    google.protobuf.Int32Value SecondMultiple = 3;
    google.protobuf.StringValue FirstAlias = 4;
    google.protobuf.StringValue SecondAlias = 5;
}

message Response{
    repeated string numbers = 1;
}
```

And the gRPC endpoint to request:

`
rpc PrintNumbers (Request) returns (Response) {}
`


To retrieve the numbers you can implement a gRPC client of your choice or use a grpCurl command, for instance:

`
grpcurl -d '{"MaxNumber":10, "FirstMultiple":3, "SecondMultiple":5, "FirstAlias":"kiss", "SecondAlias":"cool"}' -plaintext localhost:5000 model.Api/PrintNumbers
`

(See instructions here to install grpcurl: https://github.com/fullstorydev/grpcurl)