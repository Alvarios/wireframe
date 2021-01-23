MOD_FILE="go.mod"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: dep  lint

dep:
	if ! [ -f ${MOD_FILE} ]; then go mod init; fi
	go get -u ./...

lint: ## Lint the files
	go get -u golang.org/x/lint/golint
	golint -set_exit_status ${PKG_LIST}