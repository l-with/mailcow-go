.SHELLFLAGS += -x -e
PWD = $(shell pwd)
UID = $(shell id -u)
GID = $(shell id -g)

all: wget api/openapi.yaml

clean:
	rm -f *.go
	rm -rf docs/

wget:
	wget -N https://raw.githubusercontent.com/l-with/mailcow-dockerized/untype-reponses/data/web/api/openapi.yaml

openapi.yml:

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
