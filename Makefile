APP_NAME := authonchain

DESTDIR ?= /opt/$(APP_NAME)

HAS_DOCKER_COMPOSE_WITH_DASH := $(shell which docker-compose)

DEBUG_PORT ?= 40000


ifdef HAS_DOCKER_COMPOSE_WITH_DASH
	DOCKER_COMPOSE=docker-compose
else
	DOCKER_COMPOSE=docker compose
endif

install:
	$(info Installing in "$(DESTDIR)"...)
	@[ ! -d "$(DESTDIR)" ] || (echo "ERROR: Install path '$(DESTDIR)' already exists!"; exit 1)
	./scripts/build.sh "$(DESTDIR)/bin/$(APP_NAME)"
	@echo "AuthOnchain has been successfully installed in \"$(DESTDIR)\".\nEnjoy!"
terminal:
	$(DOCKER_COMPOSE) exec -u root $(APP_NAME) bash
mariadb:
	$(DOCKER_COMPOSE) exec mariadb mariadb -uroot -pauthonchain authonchain
dep-js:
	(cd frontend && npm ci --no-update-notifier --no-audit)
build-js:
	(cd frontend && npm run build)
build-go:
	rm -f $(APP_NAME)
	./scripts/build.sh $(APP_NAME)
watch-js:
	(cd frontend &&	env BUILD_ENV=development NODE_ENV=production npm run watch)
docker-all: docker-dev docker-prod
docker-prod:
	scripts/docker/build.sh authonchain
docker-dev:
	docker pull golang:1.24.1-alpine3.21
	docker pull node:22-alpine3.21
	scripts/docker/build.sh develop
debug-go:
	dlv --listen=:$(DEBUG_PORT) \
	--headless=true \
	--log=true \
	--log-output=debugger,debuglineerr,gdbwire,lldbout,rpc \
	--accept-multiclient \
	--api-version=2 \
	exec ./$(APP_NAME) -- start
