BINARY_NAME = fancurson
MAIN_PACKAGE_PATH = ./cmd

PATH_TO_PROTO_FILE = ./proto/api/order.proto
PATH_TO_PROTO = ./proto
PROTO_FOLDER_ENDPOINT = ./pkg/api/test
GOOGLEAPIS_FOLDER = ./googleapis  # Путь к googleapis

# Генерация кода для gRPC и grpc-gateway из proto-файла
.PHONY: protoc
protoc:
	protoc -I ${PATH_TO_PROTO} \
		--go_out ${PROTO_FOLDER_ENDPOINT} --go_opt paths=source_relative \
		--go-grpc_out ${PROTO_FOLDER_ENDPOINT} --go-grpc_opt paths=source_relative \
		${PATH_TO_PROTO_FILE}

.PHONY: gateway
gateway:
	protoc -I ${PATH_TO_PROTO} \
		--go_out ${PROTO_FOLDER_ENDPOINT} --go_opt paths=source_relative \
		--go-grpc_out ${PROTO_FOLDER_ENDPOINT} --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ${PROTO_FOLDER_ENDPOINT} --grpc-gateway_opt paths=source_relative \
		${PATH_TO_PROTO_FILE}

.PHONY: build
build: 
	go build -o=./tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

.PHONY: run	
run: build
	./tmp/bin/${BINARY_NAME}

.PHONY: tidy
tidy:
	go mod tidy