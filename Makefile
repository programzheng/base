IMAGE=base

.PHONY: dev, init, down
dev:
	docker-compose build web
	docker-compose up web
init:
	docker-compose build --force-rm --no-cache
	docker-compose up -d mysql
	docker-compose up -d web adminer
down:
	docker-compose down