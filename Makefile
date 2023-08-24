build:
	docker-compose build

build-pub:
	docker build . -f Dockerfile -t artarts36/service-navigator:${VERSION} --arg APP_VERSION=${VERSION} --arg BUILD_TIME=$(shell date +"%Y-%m-%dT%H:%M:%S%z")

push-pub:
	docker push artarts36/service-navigator:${VERSION}

run-f:
	docker-compose up
