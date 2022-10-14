.PHONY: all test build

GO			= go
GO_RUN		= $(GO) run
GO_BUILD	= $(GO) build
GO_TEST		= $(GO) test
GO_CLEAN	= $(GO) clean
GO_TOOL		= $(GO) tool
GO_VET		= $(GO) vet
GO_FMT		= $(GO) fmt
GO_GENERATE	= $(GO) generate
GOLINT		= golint

MAIN_GO		= cmd/main.go
BIN_NAME	= bin/gocon2022
OUT_DIR		= out/
COVER_DIR	= $(OUT_DIR)cover/
COVER_FILE	= $(COVER_DIR)cover.out
COVER_HTML	= $(COVER_DIR)cover.html

DOCKER      = docker
DOCKER_COMPOSE = $(DOCKER) compose

MIGRATE		= sql-migrate

VERSION?=0.0.0

all: clean test build

## Set up
prepare:
	@cp githooks/* .git/hooks/
	@chmod +x .git/hooks/*

## Build
gen:
	$(GO_GENERATE) ./...
run:
	$(GO_RUN) $(MAIN_GO)
build:
	$(GO_BUILD) -o $(OUT_DIR)$(BIN_NAME) $(MAIN_GO)
clean:
	$(GO_CLEAN)
	@rm -rf $(OUT_DIR)

## Test
test:
	@mkdir -p $(COVER_DIR)
	$(GO_TEST) -cover ./... -coverprofile=$(COVER_FILE)
	$(GO_TOOL) cover -html=$(COVER_FILE) -o $(COVER_HTML)

## Lint
lint:
	$(GOLINT) -set_exit_status ./...
vet:
	$(GO_VET) ./...
fmt:
	$(GO_FMT) ./...

## Docker
docker-build:
	$(DOCKER_COMPOSE) build --no-cache --force-rm #--progress=plain
up:
	$(DOCKER_COMPOSE) up -d
down:
	$(DOCKER_COMPOSE) down --remove-orphans
restart:
	$(DOCKER_COMPOSE) restart
app:
	$(DOCKER_COMPOSE) exec app sh
db:
	$(DOCKER_COMPOSE) exec db sh
log-app:
	$(DOCKER_COMPOSE) logs -f app
log-db:
	$(DOCKER_COMPOSE) logs -f db
sql:
	$(DOCKER_COMPOSE) exec db mysql -u $$MYSQL_USER -p $$MYSQL_PASSWORD $$MYSQL_DATABASE

## Migration
migrate-up:
	$(MIGRATE) up -config=configs/sql-migrate/dbconfig.yml -env=development
	$(MIGRATE) up -config=configs/sql-migrate/dbconfig.yml -env=unittest
migrate-redo:
	$(MIGRATE) redo -config=configs/sql-migrate/dbconfig.yml -env=development
	$(MIGRATE) redo -config=configs/sql-migrate/dbconfig.yml -env=unittest
migrate-down:
	$(MIGRATE) down -config=configs/sql-migrate/dbconfig.yml -env=development
	$(MIGRATE) down -config=configs/sql-migrate/dbconfig.yml -env=unittest

