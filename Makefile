# Docker image name
IMAGE_NAME := analytics-service

.PHONY: build
build:
	docker-compose build

.PHONY: run
run:
	docker run -it --rm --name $(IMAGE_NAME) -p 5555:5555 $(IMAGE_NAME)

.PHONY: hot-reload
hot-reload:
	docker run -it --rm --name $(IMAGE_NAME) -v $(PWD):/app -w /app -p 5555:5555 $(IMAGE_NAME) air

.PHONY: dev
dev:
	docker-compose up
