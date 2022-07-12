.SHELLFLAGS += -x -e
PWD = $(shell pwd)
UID = $(shell id -u)
GID = $(shell id -g)

all: clean build

clean:
	rm -f *.go
	rm -rf docs/

build:
	wget -O openapi.yml https://raw.githubusercontent.com/l-with/mailcow-dockerized/openapi-3.0.3/data/web/api/openapi.yaml
	docker run \
		--rm -v ${PWD}:/local \
		--user ${UID}:${GID} \
		openapitools/openapi-generator-cli:v6.0.1 generate \
		-i /local/openapi.yml \
		-g go \
		-o /local \
		-c /local/config.yaml
	go get
	go fmt .
	go mod tidy
	rm -f .travis.yml
