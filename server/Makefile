@all:

api-public-dev:
	air -c ./cmd/api-public/.air.toml

api-public:
	docker build . -f ./Dockerfile.api-public --platform=linux/amd64 -t ghcr.io/codeserk/media-api-public:$(shell git rev-parse HEAD) --push
api-public-latest:
	docker build . -f ./Dockerfile.api-public --platform=linux/amd64 -t ghcr.io/codeserk/media-server-api-public:latest --push
