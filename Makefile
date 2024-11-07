project ?= lottery
host ?= 127.0.0.1:3306
user ?= root
pwd ?=
table ?=
cache ?=

init:
	go install github.com/zeromicro/go-zero/tools/goctl@latest
	goctl env check --install --verbose --force
	go install github.com/zeromicro/goctl-swagger@latest

api-format:
	goctl api format --dir ./doc/api/
	goctl api plugin -plugin goctl-swagger="swagger -filename ${project}.json -basepath /api" -api ./doc/api/${project}.api -dir ./doc/swagger

api:
	make api-format
	goctl api go -api ./doc/api/${project}.api --dir . --style go_zero

db:
	goctl model mysql datasource -url="${user}:${pwd}@tcp(${host})/${table}" -table="${table}" --dir internal/model --strict --style go_zero ${cache}
