APP_NAME := authonchain

DESTDIR ?= /opt/$(APP_NAME)

HAS_DOCKER_COMPOSE_WITH_DASH := $(shell which docker-compose)


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
build-go:
	rm -f $(APP_NAME)
	./scripts/build.sh $(APP_NAME)
docker-all: docker-dev docker-prod
docker-prod:
	scripts/docker/build.sh authonchain
docker-dev:
	docker pull golang:1.24.1-alpine3.21
	scripts/docker/build.sh develop
