#goctl rpc protoc ./rpc/user.proto --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc    #grpc生成命令

#goctl usermodel mysql ddl -src="../../deploy/sql/user.sql" -dir="./usermodel" -c


#goctl api go -api ./api/user.api -dir apps/user/api -style gozero