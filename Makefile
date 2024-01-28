DATE=$(shell date +%s)

docker-build:
	docker build . --file Dockerfile --tag heroyt/lac-event-server:latest --tag heroyt/lac-event-server:${DATE}

docker-build-amd:
	docker build . --file Dockerfile --platform linux/amd64 --tag heroyt/lac-event-server:latest --tag heroyt/lac-event-server:${DATE}

docker-push:
	docker push heroyt/lac-event-server -a