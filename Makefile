DOCKER=docker-compose

up:
	${DOCKER} up -d
up-build:
	${DOCKER} up --build -d
down:
	${DOCKER} down -v
build:
	${DOCKER} build
exec:
	${DOCKER} exec api go run main.go