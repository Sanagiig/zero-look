# proto file path: ./app/usercenter/cmd/rpc/pb/usercenter.proto

goctl rpc protoc ./app/$1/cmd/rpc/pb/$1.proto --go_out=./app/$1/cmd/rpc --go-grpc_out=./app/$1/cmd/rpc --zrpc_out=./app/$1/cmd/rpc/ --home ./develop/.goctl --style go_zero