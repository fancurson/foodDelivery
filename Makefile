PROTO_FILE_PATH = ./api/order.proto
PROTO_FOLDER_ENDPOINT = ./pkg/api/test
GOOGLEAPIS_FOLDER = ./googleapis  # Путь к googleapis

# Генерация кода для gRPC и grpc-gateway из proto-файла
.PHONY: protoc
protoc:
	protoc -I ./api/ -I ${GOOGLEAPIS_FOLDER} \
   --go_out ${PROTO_FOLDER_ENDPOINT} --go_opt paths=source_relative \
   --go-grpc_out ${PROTO_FOLDER_ENDPOINT} --go-grpc_opt paths=source_relative \
   ${PROTO_FILE_PATH}

.PHONY: geteway
geteway: 
   protoc -I ./api/ -I ./googleapis 
   --go_out ./pkg/api/test --go_opt paths=source_relative \
   --go-grpc_out ./pkg/api/test --go-grpc_opt paths=source_relative \
   --grpc-gateway_out ./pkg/api/test --grpc-gateway_opt paths=source_relative \
   ./api/order.proto