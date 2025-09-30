SHELL=/bin/bash

.PHONY: test
test:
	@pushd fizzbuzz > /dev/null && go test ./... && popd > /dev/null

.PHONY: vet
vet:
	@pushd fizzbuzz > /dev/null && go vet ./... && popd > /dev/null

.PHONY: format
format:
	@pushd fizzbuzz > /dev/null && go fmt ./... && popd > /dev/null

.PHONY: deps
deps:
	@pushd fizzbuzz > /dev/null && go mod download && go mod verify && popd > /dev/null

.PHONY: pre-deploy
pre-deploy: deps format vet test

.PHONY: logs
logs:
	docker compose logs -f

.PHONY: dev-docker
dev-docker:
	docker compose up --build -d

.PHONY: down
down:
	docker compose down

.PHONY: swagg.fmt
swagg.fmt:
	@pushd fizzbuzz > /dev/null && go run github.com/swaggo/swag/cmd/swag@latest fmt -g cmd/server/main.go && popd > /dev/null

.PHONY: swagg
swagg: swagg.fmt
	@pushd fizzbuzz > /dev/null && go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/server/main.go --parseDependency --parseInternal && popd > /dev/null