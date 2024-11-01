package config

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"net/url"
)

type Config struct {
	rest.RestConf
	Mysql MySqlConf
	Cache cache.CacheConf
}

type MySqlConf struct {
	Addr     string
	User     string
	Password string
	Database string
	Loc      string `json:",default=Local"`
}

func (m MySqlConf) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=%s", m.User, m.Password, m.Addr, m.Database, url.QueryEscape(m.Loc))
}

func (m MySqlConf) Conn(opts ...sqlx.SqlOption) sqlx.SqlConn {
	return sqlx.NewMysql(m.Dsn(), opts...)
}
