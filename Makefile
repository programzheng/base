IMAGE=base
all: dev, down
.PHONY: all
dev:
	docker-compose build --force-rm
	docker-compose up -d mysql
	docker-compose up -d web adminer
down:
	docker-compose down