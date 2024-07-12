BUILD_PATH=bin/main
MAIN_PACKAGE_PATH=cmd/main.go
default: hot_reload

## help: Print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## run: Run the application from the build
.PHONY: run
run: build
	@if [ "$(shell uname)" = "Linux" ]; then \
	echo "Running Linux build"; \
		./${BUILD_PATH}-linux; \
	elif [ "$(shell uname)" = "Windows" ]; then \
	echo "Running Windows build"; \
		./${BUILD_PATH}-windows; \
	else \
		echo "Unsupported operating system"; \
	fi

## run/live: Run the application with reloading on file changes
.PHONY: run/live
hot_reload:
	@nodemon --exec go run ${MAIN_PACKAGE_PATH} --signal SIGTERM

## build: Build the application
.PHONY: build
build:
	@echo "Start building\n"
	@GOARCH=amd64 GOOS=linux go build -o ${BUILD_PATH}-linux ${MAIN_PACKAGE_PATH}
	@echo "linux build complete\n"
	@GOARCH=amd64 GOOS=windows go build -o ${BUILD_PATH}-windows ${MAIN_PACKAGE_PATH}
	@echo "windows build complete\n"
	@echo "Build complete\n"

## test: Run the tests
.PHONY: test
test:
	@go test -v ./...

## format: Format all go files in the repository
.PHONY: format
format:
	@go fmt ./...
	@golines -w -m 120 .
