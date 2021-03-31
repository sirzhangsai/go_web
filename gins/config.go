package gins

import (
	"fmt"
)
type Config struct {

	Port  int
	Host  string
}
//服务ip配置
func(conf *Config) Addr() string {

	return fmt.Sprintf("%s:%d", conf.Host, conf.Port)
}