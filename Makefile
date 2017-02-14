.DEFAULT_GOAL := help

IMAGE_NAME_TEST := bacon-test
IMAGE_NAME_PROD := bacon-prod

.PHONY: help build test clean

help:
		@echo "------------------------------------------------------------------------"
		@echo Bacon
		@echo "------------------------------------------------------------------------"
		@grep -E '^[a-zA-Z_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: clean ## Build unit test container
		@docker build -t ${IMAGE_NAME_TEST} -f resources/test/Dockerfile .

test: build ## Run unit tests
		@docker run --rm ${IMAGE_NAME_TEST}

clean: ## Remove images and containers
		@./resources/scripts/rm-image.sh ${IMAGE_NAME_TEST}
		@./resources/scripts/rm-container.sh ${IMAGE_NAME_TEST}

run: ## Start container
		@./resources/scripts/rm-container.sh ${IMAGE_NAME_PROD}
		@docker build -t ${IMAGE_NAME_PROD} -f resources/prod/Dockerfile .
		@docker run -d -p 8088:8088 --name bacon-prod ${IMAGE_NAME_PROD}
