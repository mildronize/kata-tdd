test:
	gotest -v ./...

watch-test:
	nodemon -e go --exec "gotest -v ./... || exit 1" 