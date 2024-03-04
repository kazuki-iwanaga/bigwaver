DC := docker compose -f compose.yaml --env-file compose.env

.PHONY: build
build:
	$(DC) build --no-cache

.PHONY: up
up:
	$(DC) up -d

.PHONY: down
down:
	$(DC) down

.PHONY: in
in:
	$(DC) exec app bash

.PHONY: ps
ps:
	$(DC) ps

.PHONY: logs
logs:
	$(DC) logs

.PHONY: test
test:
	curl -X POST -H "Content-Type: application/json" -d '{"header1":"header1","data":{"header2":"header2"}}' localhost:8080/hooks/github