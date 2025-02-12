.PHONY: vendor

# Main
# ============================================================================

help: ## This help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Docker
# ============================================================================

export HOST_UID ?= $(shell id -u)
export HOST_GID ?= $(shell id -g)

IMAGES := $(shell docker images -q "boilerplate-*")
CONTAINERS := $(shell docker ps -aq --filter name=boilerplate-)

up: ## Up docker application
	mkdir -p docker/var/postgres
	docker-compose -f docker-compose.yml up -d --remove-orphans

up-local: ## Up local application
	mkdir -p docker/var/postgres
	docker-compose -f docker-compose.local.yml up -d --remove-orphans

down: ## Down application
ifdef CONTAINERS
	docker stop $(CONTAINERS)
endif

clean: down ## Remove containers
ifdef CONTAINERS
	docker rm -f --volumes $(CONTAINERS)
endif

clean-all: clean ## Remove containers, images and networks
ifdef IMAGES
	docker rmi -f $(IMAGES)
endif
	docker network rm -f boilerplate-network

logs: ## Logs
	docker-compose logs --follow
#	docker logs --follow $(CONTAINERS)

# Docker shell
# ============================================================================

shell-app: ## Shell of postgresql container
	docker exec --interactive --tty boilerplate-app /bin/bash

shell-postgres: ## Shell of postgresql container
	docker exec --interactive --tty boilerplate-postgres /bin/bash

# Modules
# ============================================================================

vendor: ## Go mod vendor
	docker exec boilerplate-app go mod tidy
	docker exec boilerplate-app go mod vendor
