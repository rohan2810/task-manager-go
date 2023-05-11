CURRENT_DIR = $(shell pwd)
export PATH := $(CURDIR)/bin:$(PATH)

api-lint:
	api-linter apis/*.proto

buf-lint:
	@echo "linting proto files"
	buf lint