DOCKER=docker-compose

up:
	${DOCKER} up -d

up build:
	${DOCKER} up --build -d

down:
	${DOCKER} down

down v:
	${DOCKER} down -v

exec:
	${DOCKER} exec api go run main.go