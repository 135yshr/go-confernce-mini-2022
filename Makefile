.PHONY: all test build

GO			= go
GO_RUN		= $(GO) run
GO_BUILD	= $(GO) build
GO_TEST		= $(GO) test
GO_CLEAN	= $(GO) clean
GO_TOOL		= $(GO) tool
GO_VET		= $(GO) vet
GOLINT		= golint

MAIN_GO		= cmd/main.go
BIN_NAME	= build/smacare
OUT_DIR		= out/
COVER_DIR	= $(OUT_DIR)cover/
COVER_FILE	= $(COVER_DIR)cover.out
COVER_HTML	= $(COVER_DIR)cover.html

VERSION?=0.0.0

all: clean test build

## Set up
prepare:
	@cp githooks/* .git/hooks/
	@chmod +x .git/hooks/*

## Build
run:
	$(GO_RUN) $(MAIN_GO)
build:
	$(GO_BUILD) -o $(BIN_DIR)$(BIN_NAME) $(MAIN_GO)
clean:
	$(GO_CLEAN)
	@rm -rf $(BIN_DIR)$(BIN_NAME)

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

