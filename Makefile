#讀取.env
include ./.env
export $(shell sed 's/=.*//' ./.env)

#當前年-月-日
DATE=$(shell date +"%F")
COMPOSE=docker-compose
BASH?=bash
WEB=api

.PHONY: dev, up, init, down
bash:
	$(COMPOSE) exec $(WEB) $(BASH)

#重新編譯
dev:
	$(COMPOSE) build $(WEB)
	$(COMPOSE) up $(WEB)

#啟動服務
up:
	$(COMPOSE) up -d mysql adminer minio
	$(MAKE) dev

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

ngrok:
	$(COMPOSE) up ngrok
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

#備份mysql all database
mysql-backup:
	$(COMPOSE) up -d mysql
	$(MAKE) check-data-directory-mysql-backup
	cp -R -f $(DATA_PATH_HOST)/mysql $(DATA_PATH_HOST)/mysql-backup/$(DATE)
	# remove 3 days ago backup directory
	rm -r $(DATA_PATH_HOST)/mysql-backup/$(shell date --date="3 days ago" +"%F")

#檢查資料夾並建立
check-data-directory-%:
	if test -d $(DATA_PATH_HOST)/$*; \
	then echo $* is exists; exit 0; \
	else mkdir $(DATA_PATH_HOST)/$*; \
	fi