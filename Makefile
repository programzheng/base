IMAGE=base

.PHONY: dev, down
dev:
	docker-compose build --force-rm
	docker-compose up -d mysql
	docker-compose up -d web adminer
down:
	docker-compose down