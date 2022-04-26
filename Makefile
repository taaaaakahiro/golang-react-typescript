DOCKER=docker-compose

up:
	${DOCKER} up --build -d

down:
	${DOCKER} down -v

exec:
	${DOCKER} exec api go run main.go