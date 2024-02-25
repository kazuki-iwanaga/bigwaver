APP=rath

.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: ps
ps:
	docker-compose ps

.PHONY: build
build:
	docker-compose build --no-cache

.PHONY: in
in:
	docker-compose exec $(APP) bash