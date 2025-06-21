.PHONY: vendor

# Main
# ============================================================================

help: ## This help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Docker
# ============================================================================

export HOST_UID ?= $(shell id -u)
export HOST_GID ?= $(shell id -g)

IMAGES := $(shell docker images -q "goapp-*")
CONTAINERS := $(shell docker ps -aq --filter name=goapp-)

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
	docker network rm -f goapp-network

test: ## Test
	docker exec goapp-app go test -v ./...

logs: ## Logs
	docker-compose logs --follow

# Postgresql
# ============================================================================

psql: ## Psql client
	docker exec --interactive --tty goapp-postgres psql -d goapp_db

test-db: ## Create or recreate empty database for e2e testing (dbname=test_db)
	@echo "drop database test_db\n"
	docker exec goapp-postgres psql -c "drop database if exists test_db"
	@echo "\ncreate database test_db\n"
	docker exec goapp-postgres psql -c "create database test_db"
	@echo ""
	go run ./cmd/cli/main.go migrate --dbname test_db

# Docker shell
# ============================================================================

shell-app: ## Shell of postgresql container
	docker exec --interactive --tty goapp-app /bin/bash

shell-postgres: ## Shell of postgresql container
	docker exec --interactive --tty goapp-postgres /bin/bash

# Modules
# ============================================================================

vendor: ## Go mod vendor
	docker exec goapp-app go mod tidy
	docker exec goapp-app go mod vendor

tools: ## Install develop tools
	docker exec goapp-app go install github.com/boumenot/gocover-cobertura@latest
	docker exec goapp-app go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	docker exec goapp-app go install github.com/jondot/goweight@latest
	docker exec goapp-app go install github.com/psampaz/go-mod-outdated@latest
	docker exec goapp-app go install github.com/rakyll/gotest@latest
	docker exec goapp-app go install github.com/sonatype-nexus-community/nancy@latest