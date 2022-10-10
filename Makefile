DOCKER=docker-compose

#　起動
run:
	${DOCKER} up -d
up-build-d:
	${DOCKER} up --build -d
#　停止
down:
	${DOCKER} down
down-v:
	${DOCKER} down -v
# 再起動
re:
	${DOCKER} restart
# ビルド
build:
	${DOCKER} build
# コンテナ内
exec:
	${DOCKER} exec api go run main.go
sh:
	${DOCKER} exec api sh
# Test
test:
	${DOCKER} exec api go test -v -cover