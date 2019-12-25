COMPOSE=docker-compose
BASH?=bash

.PHONY: dev, up, init, down
dev:
	$(COMPOSE) build api
	$(COMPOSE) up api
up:
	$(COMPOSE) up -d mysql adminer minio
	$(MAKE) dev
bash:
	$(COMPOSE) exec api $(BASH)
init:
	$(COMPOSE) build --force-rm --no-cache
	$(MAKE) up
ps:
	$(COMPOSE) ps
down:
	$(COMPOSE) down