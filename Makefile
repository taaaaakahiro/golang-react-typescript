DOCKER=docker-compose

#　起動
up:
	${DOCKER} up
up-d:
	${DOCKER} up -d
up-build-d:
	${DOCKER} up --build -d
#　停止
down:
	${DOCKER} down
down-v:
	${DOCKER} down -v
# 再起動
restart:
	${DOCKER} restart
# ビルド
build:
	${DOCKER} build
# コンテナ内
exec:
	${DOCKER} exec api go run main.go
sh:
	${DOCKER} exec api sh