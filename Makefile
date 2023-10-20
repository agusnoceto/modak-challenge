APP       ?= modak/challenge
BUILD_TAG ?= local
WORKDIR   ?= $(shell pwd)
BINDIR    ?= $(WORKDIR)/bin
GO        ?= $(shell which go)
DOCKER    ?= $(shell which docker)

.PHONY: help
help: ## Show this help.
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: check-deps
check-deps: ## Check if docker is installed.
ifeq ($(strip $(DOCKER)),)
	@echo "Docker is not available. Please install docker."
	@exit 1
endif

.PHONY: build
build: ## Build locally.
	$(GO) build -o $(BINDIR)/$(APP) ./cmd

.PHONY: clean
clean: ## Clean local build directory.
	@$(RM) -r $(BINDIR)

.PHONY: test
test: ## Run tests.
	$(GO) test -v ./...

.PHONY: run
run: ## Run this locally.
	$(BINDIR)/$(APP)

.PHONY: build-docker
build-docker: check-deps ## Build inside docker container.
	$(DOCKER) build --build-arg GIT_CREDS=$$GIT_CREDS --tag $(APP):$(BUILD_TAG) -f Dockerfile .

.PHONY: run-docker
run-docker: build-docker ## Run the docker image.
	$(DOCKER) run -it $(APP):$(BUILD_TAG)