# go-zero-lottery

This is a [go-zero](https://go-zero.dev) project. Lottery hero for HOK. And same logic project for rust [actix-web](https://github.com/weiraneve/hok-lottery-actix)

I have kotlin SpringBoot project as some logic [hok-lottery](https://github.com/weiraneve/hok-lottery)

## apis

See [the API documentation](./doc/api) for more info.

## sql

See [lottery sql info](./doc/sql/lottery.sql) for lottery sql info

## usages

When you config your mysql config, then:
```
go run lottery.go -f etc/lottery.yaml
```

When you update .api file, then you can run`goctl api go -api pick.api -dir .`