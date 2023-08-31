GOPATH:=$(shell go env GOPATH)
APP?=server
ifneq (,$(wildcard ./.env))
    include .env
    export
endif
IMAGE_NAME?="${REGION}-docker.pkg.dev/${PROJECT_ID}/${REPOSITORY_NAME}/${APP}"

.PHONY: init
## init: initialize the application(server)
init:
	go mod download

.PHONY: build
## build: build the application(server)
build:
	go build -o build/${APP} cmd/main.go

.PHONY: run
## run: run the application(server)
run:
	go run -v -race cmd/main.go

.PHONY: format
## format: format files
format:
	gofmt -s -w .
	go mod tidy

.PHONY: test
## test: run tests
test:
	@go install github.com/rakyll/gotest@latest
	gotest -race -cover -v ./...

.PHONY: coverage
## coverage: run tests with coverage
coverage:
	@go install github.com/rakyll/gotest@latest
	gotest -race -coverprofile=coverage.txt -covermode=atomic -v ./...

.PHONY: lint
## lint: check everything's okay
lint:
	golangci-lint run ./...
	go mod verify

.PHONY: generate
## generate: generate source code for mocking
generate:
	@go install golang.org/x/tools/cmd/stringer@latest
	@go install github.com/golang/mock/mockgen@latest
	go generate ./...
	@$(MAKE) format

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':'

.PHONY: deploy-to-cloud-run
## deploy-to-cloud-run: deploy to cloud run
deploy-to-cloud-run:
	docker build -t ${IMAGE_NAME} . && \
	docker push ${IMAGE_NAME} && \
	gcloud run deploy server --image=${IMAGE_NAME} --platform managed --region ${REGION} --allow-unauthenticated --project ${PROJECT_ID} --set-env-vars "KAKAO_REST_API_KEY=${KAKAO_REST_API_KEY}"
