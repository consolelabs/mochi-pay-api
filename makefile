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

dependencies:
	go install github.com/rubenv/sql-migrate/...@latest
	go install -v github.com/golang/mock/mockgen@v1.6.0
	go install -v github.com/vektra/mockery/v2@v2.15.0
	go install -v github.com/swaggo/swag/cmd/swag@v1.8.7

remove-infras:
	docker-compose down --remove-orphans
	
init:
	docker-compose up -d
	@echo "Waiting for database connection..."
	@while ! docker exec mochi_pay_api_local pg_isready > /dev/null; do \
		sleep 1; \
	done