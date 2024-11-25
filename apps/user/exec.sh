#goctl rpc protoc ./rpc/user.proto --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc    #grpc生成命令

goctl model mysql ddl -src="../../deploy/sql/user.sql" -dir="./model" -c