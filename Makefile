DC := docker compose

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
	$(DC) exec app sh

.PHONY: ps
ps:
	$(DC) ps