BINARY_NAME=social_network

build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}

run:
	./${BINARY_NAME}

test:
	go test

lint:
	golangci-lint run

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}-linux