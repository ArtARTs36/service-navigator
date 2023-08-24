build:
	docker-compose build

build-pub:
	docker build . -f Dockerfile -t artarts36/service-navigator

run-f:
	docker-compose up
