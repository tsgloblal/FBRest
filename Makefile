SHELL=/bin/bash

.PHONY: dev-docker
dev-docker:
	docker compose up --build -d

.PHONY: swagg.fmt
swagg.fmt:
	@pushd fizzbuzz > /dev/null && go run github.com/swaggo/swag/cmd/swag@latest fmt -g cmd/server/main.go && popd > /dev/null

.PHONY: swagg
swagg: swagg.fmt
	@pushd fizzbuzz > /dev/null && go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/server/main.go --parseDependency --parseInternal && popd > /dev/null