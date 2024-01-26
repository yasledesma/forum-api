APP_NAME=forum_api

init: # Start docker container
	@sudo docker build -t ${APP_NAME} .
	@sudo docker run --rm -d -p 8080:8080 -v ${PWD}:/usr/src/build -it --name ${APP_NAME} ${APP_NAME} 

stop: # Stop docker container
	@sudo docker stop ${APP_NAME}

term: # Spin up an interactive terminal inside the container
	@sudo docker exec -it ${APP_NAME} sh

clean: # Remove build and temporary files
	@sudo docker exec ${APP_NAME} go clean
	@sudo docker exec ${APP_NAME} rm -rf ./out/* ./tmp/*.log

vendor: # Copy all dependencies to /vendor
	@sudo docker exec ${APP_NAME} go mod vendor

test: # Run all test suites
	@sudo docker exec ${APP_NAME} go test -v ./...

coverage: # Run tests coverage
	@sudo docker exec ${APP_NAME} mkdir -p ./out
	@sudo docker exec ${APP_NAME} go test -cover -covermode=count -coverprofile=./out/coverage.out ./...
	@sudo docker exec ${APP_NAME} go tool cover -func ./out/coverage.out 
	@sudo docker exec ${APP_NAME} go tool cover -html=./out/coverage.out -o ./out/coverage.html
	
lint: # Lint project
	@sudo docker exec ${APP_NAME} golangci-lint run

# Non-container commands
build: # Build binary
	@mkdir -p ./out/bin
	@GO111MODULE=on go build -mod vendor -o ./out/bin/${APP_NAME} .

run: build # Build and run binary
	@./out/bin/${BINARY_NAME}

