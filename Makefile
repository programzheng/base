COMPOSE=docker-compose
BASH?=bash

.PHONY: dev, up, init, down
dev:
	$(COMPOSE) build web
	$(COMPOSE) up web
up:
	$(COMPOSE) up -d mysql adminer minio
	$(MAKE) dev
bash:
	$(COMPOSE) exec web $(BASH)
init:
	$(COMPOSE) build --force-rm --no-cache
	$(MAKE) up
ps:
	$(COMPOSE) ps
down:
	$(COMPOSE) down