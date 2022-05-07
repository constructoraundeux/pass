.PHONY: build
build:
	@echo 'Building for linux'
	GOOS=linux GOARCH=amd64 go build -o=./bin/linux/undeux .
	@echo 'Building for windows'
	GOOS=windows GOARCH=amd64 go build -o=./bin/windows/undeux.exe .

NEW_PATCH=$(shell git describe --tags --abbrev=0 | tr -d "v" | awk -F '.' '{printf "v"$$1"."$$2"."$$3+1}')
NEW_MINOR=$(shell git describe --tags --abbrev=0 | tr -d "v" | awk -F '.' '{printf "v"$$1"."$$2+1"."$$3}')
NEW_MAJOR=$(shell git describe --tags --abbrev=0 | tr -d "v" | awk -F '.' '{printf "v"$$1+1"."$$2"."$$3}')

.PHONY: release/patch
release/patch:
	@echo ${NEW_PATCH}
	sed -i -E "s/v[0-9]*\.[0-9]*\.[0-9]*/${NEW_PATCH}/g" release.json
	git tag ${NEW_PATCH}
	git commit -am 'Set version ${NEW_PATCH}'
	git push origin main
	sleep 3
	git push --tags

.PHONY: release/minor
release/minor:
	@echo ${NEW_MINOR}
	sed -i -E "s/v[0-9]*\.[0-9]*\.[0-9]*/${NEW_MINOR}/g" release.json
	git tag ${NEW_MINOR}
	git commit -am 'Set version ${NEW_MINOR}'
	git push origin main
	sleep 3
	git push --tags

.PHONY: release/major
release/major:
	@echo ${NEW_MAJOR}
	sed -i -E "s/v[0-9]*\.[0-9]*\.[0-9]*/${NEW_MAJOR}/g" release.json
	git tag ${NEW_MAJOR}
	git commit -am 'Set version ${NEW_MAJOR}'
	git push origin main
	sleep 3
	git push --tags

