.SHELLFLAGS += -x -e
PWD = $(shell pwd)
UID = $(shell id -u)
GID = $(shell id -g)

all: clean build

clean:
	rm -f *.go
	rm -rf docs/

build:
	wget -O schema.yml https://raw.githubusercontent.com/l-with/mailcow-dockerized/improve-domain-api-schema/data/web/api/openapi.yaml
	docker run \
		--rm -v ${PWD}:/local \
		--user ${UID}:${GID} \
		openapitools/openapi-generator-cli:v6.0.0 generate \
		-i /local/schema.yml \
		-g go \
		-o /local \
		-c /local/config.yaml
	go get
	go fmt .
	go mod tidy
	rm -f .travis.yml
