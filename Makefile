BINARY_NAME=forum_api

build: # Build binary
	mkdir -p ./out/bin
	GO111MODULE=on go build -mod vendor -o ./out/bin/$(BINARY_NAME) .

run: build # Build and run binary
	./out/bin/${BINARY_NAME}

watch: # Hot reload on change
	${GOPATH}/air -c .air.toml

clean: # Remove build and temporary files
	go clean
	rm -rf ./out ./tmp

vendor: # Copy all dependencies to /vendor
	go mod vendor

test: # Run all test suites
	go test -v ./...

test_coverage: # Run tests coverage
	mkdir -p ./out
	go test -cover -covermode=count -coverprofile=./out/coverage.out ./...
	go tool cover -func ./out/coverage.out 
	go tool cover -html=./out/coverage.out -o ./out/coverage.html
	
# Saving this for when CI is implemented in this project
# lint: # Lint project # https://golangci-lint.run/usage/install/
# 	golang-ci-lint run --enable-all

