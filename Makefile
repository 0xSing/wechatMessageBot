.PHONY: walletSynV2

GO := go
BUILD_DIR := ./build
BIN_DIR := $(BUILD_DIR)/bin
CMD_DIR := ./
MAIN_PACKAGE := main.go
SERVER_NAME := wechatBot	# config Server name

wechatBot:
	$(GO) build -o $(BIN_DIR)/$(SERVER_NAME) $(CMD_DIR)$(MAIN_PACKAGE)
	@echo "Done building."
	@echo "Run \"$(BIN_DIR)/$(SERVER_NAME)\" to launch $(SERVER_NAME)."
