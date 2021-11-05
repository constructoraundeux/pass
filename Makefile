.PHONY: build
build:
	@echo 'Building for linux'
	GOOS=linux GOARCH=amd64 go build -o=./bin/linux/undeux .
	@echo 'Building for windows'
	GOOS=windows GOARCH=amd64 go build -o=./bin/windows/undeux.exe .

