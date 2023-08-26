build:
	docker-compose build

build-pub:
	docker build . -f Dockerfile -t artarts36/service-navigator:${VERSION} --build-arg APP_VERSION=${VERSION} --build-arg BUILD_TIME=$(shell date +"%Y-%m-%dT%H:%M:%S%z")

push-pub:
	docker push artarts36/service-navigator:${VERSION}

run-f:
	docker-compose up

lint:
	docker run \
        --rm -t \
        -v "${PWD}/:/app" \
        --workdir /app \
        golangci/golangci-lint \
        golangci-lint run

test:
	go test ./...
