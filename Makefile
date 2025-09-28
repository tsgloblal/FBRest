.PHONY: dev-docker
dev-docker:
	docker compose up --build -d

.PHONY: build
build:
	go build -o bin/main .

.PHONY: run
run:
	./bin/main