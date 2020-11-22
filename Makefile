#讀取.env
include ./.env
export $(shell sed 's/=.*//' ./.env)

#當前年-月-日
DATE=$(shell date +"%F")
COMPOSE=docker-compose
BASH?=bash
SERVICES=api

.PHONY: dev, up, init, down
bash:
	$(COMPOSE) exec $(SERVICES) $(BASH)

#重新編譯
dev:
	$(COMPOSE) build $(SERVICES)
	$(COMPOSE) up $(SERVICES)

#啟動服務
up:
	$(COMPOSE) up -d $(SERVICES)

#重啟服務
restart:
	$(COMPOSE) restart

#初始化
init:
	$(COMPOSE) build --force-rm --no-cache
	$(MAKE) up
#列出容器列表
ps:
	$(COMPOSE) ps

#服務log
#%=service name
logs-%:
	$(COMPOSE) logs $*

#關閉所有服務
down:
	$(COMPOSE) down

#移除多餘的image
prune:
	docker system prune