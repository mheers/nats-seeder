all: build

build:
	docker build --no-cache=true -t mheers/nats-seeder .

push:
	docker push mheers/nats-seeder

start:
	docker-compose up
