MAIN_DIR = "cmd/api/main.go"
ROUTER_DIR = "cmd/pkg/api/routers/router.go"
TARGET = "cmd/api/"
BINARY_NAME = "snap_roles"

build:
	go build -o ${BINARY_NAME} ${MAIN_DIR}

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

swag:
	swag init -g ${ROUTER_DIR}

swag_run: swag run
