IMAGE=base

.PHONY: dev, up, init, down
dev:
	docker-compose build web
	docker-compose up web
up:
	docker-compose up -d mysql
	docker-compose up -d web adminer
init:
	docker-compose build --force-rm --no-cache
	$(MAKE) up
down:
	docker-compose down