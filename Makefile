.SHELLFLAGS += -x -e
PWD = $(shell pwd)
UID = $(shell id -u)
GID = $(shell id -g)

all: api/openapi.yaml

clean:
	rm -f *.go
	rm -rf docs/

openapi.yaml:
	wget -N https://raw.githubusercontent.com/l-with/mailcow-dockerized/improve-domain-api-schema/data/web/api/openapi.yaml

api/openapi.yaml: openapi.yaml
	rm -f *.go
	rm -rf docs/
	docker run \
		--rm -v ${PWD}:/local \
		--user ${UID}:${GID} \
		openapitools/openapi-generator-cli:v6.0.0 generate \
		-i /local/openapi.yaml \
		-g go \
		-o /local \
		-c /local/config.yaml
	go get
	go fmt .
	go mod tidy
	rm -f .travis.yml

build-internal:
	rm -f *.go
	rm -rf docs/
	docker run \
		--rm -v ${PWD}:/local \
		--user ${UID}:${GID} \
		openapitools/openapi-generator-cli:v6.0.0 generate \
		-i /local/openapi.yaml \
		-g go \
		-o /local \
		-c /local/config.yaml
	go get
	go fmt .
	go mod tidy
	rm -f .travis.yml
