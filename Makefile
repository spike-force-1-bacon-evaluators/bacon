.DEFAULT_GOAL := help

IMAGE_NAME := bacon

.PHONY: help build test clean

help:
	@echo "------------------------------------------------------------------------"
	@echo Bacon
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: clean ## Build unit test container
	@docker build -t ${IMAGE_NAME} -f resources/test/Dockerfile .

test: build ## Run unit tests
	@docker run --rm ${IMAGE_NAME}

clean: ## Remove images and containers
	@./resources/scripts/rm-image.sh ${IMAGE_NAME}
	@./resources/scripts/rm-container.sh ${IMAGE_NAME}
