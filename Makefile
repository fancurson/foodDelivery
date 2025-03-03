PROTO_FILE_PATH =  ./api/order.proto
PROTO_FOLDER_ENDPOINT = ./pkg/api/test


# Генерация кода для gRPC и grpc-gateway из proto-файла
.PHONY: protoc
protoc:
	protoc -I ./api \
   --go_out ${PROTO_FOLDER_ENDPOINT} --go_opt paths=source_relative \
   --go-grpc_out ${PROTO_FOLDER_ENDPOINT} --go-grpc_opt paths=source_relative \
   ${PROTO_FILE_PATH}