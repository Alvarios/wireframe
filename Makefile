
MOD_FILE="go.mod"

dep:
	if ! [ -f ${MOD_FILE} ]; then go mod init; fi
	go get -u ./...