DATE=$(shell date +%s)

docker:
	docker buildx build --platform linux/arm64,linux/amd64 . --file Dockerfile --tag heroyt/lac-event-server:latest --tag heroyt/lac-event-server:${DATE} --push

docker-build:
	docker build . --file Dockerfile --tag heroyt/lac-event-server:latest --tag heroyt/lac-event-server:${DATE}

docker-build-amd:
	docker build . --file Dockerfile --platform linux/amd64 --tag heroyt/lac-event-server:latest --tag heroyt/lac-event-server:${DATE}

docker-push:
	docker push heroyt/lac-event-server -a