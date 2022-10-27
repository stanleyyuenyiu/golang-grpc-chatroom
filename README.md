# golang-grpc-chatroom

## Introduction
A poc to use gRPC streaming as chat room

## Prerequisite
- Install protobuf compliter 3.x -> https://github.com/protocolbuffers/protobuf/ (not require for client)
- Install golang 1.19+
- Install ngrok (only require if test on remote client)

## How to Run
```
// Start a grpc server
go run server/main.go -port {your port}

// Expose server to public if test run on remote
ngrok tcp {your port}

// Start client
go run client/main.go -user {your userid} -name {your username} -channel {chatroom channel} -addr {your server port or remote server address}
```
