APP_NAME=mochi-pay-api
TOOLS_IMAGE=huynguyenh/console-tools
APP_ENVIRONMENT=docker run --rm -v ${PWD}:/${APP_NAME} -w /${APP_NAME} --net=host ${TOOLS_IMAGE}

dev:
	go run cmd/api/api.go

test:
	go test ./...

build-tooling:
	docker build -f ./Dockerfile.tools -t ${TOOLS_IMAGE} .

push-tooling:
	docker push ${TOOLS_IMAGE}

tooling:
	docker pull ${TOOLS_IMAGE}

migrate-up:
	${APP_ENVIRONMENT} sql-migrate up -env="local"

migrate-new:
	${APP_ENVIRONMENT} sql-migrate new -env=local ${name}
