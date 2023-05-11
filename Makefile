CURRENT_DIR = $(shell pwd)
export PATH := $(CURDIR)/bin:$(PATH)

api-lint:
	api-linter apis/v1/*.proto

buf-lint:
	@echo "linting proto files"
	buf lint

generate-go:
	protoc --go_out=. apis/v1/*.proto