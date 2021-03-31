package config
import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)
type Configs struct {
	DBs []struct {
		Alias string		`yaml:"alias"`
		Type  string		`yaml:"type"`
		Server string		`yaml:"server"`
		Port   int		`yaml:"port"`
		Database string		`yaml:"database"`
		User string			`yaml:"user"`
		Password string		`yaml:"password"`
	} `yaml:"dbs"`
	Cache struct {
		Redis Cache  `yaml:"redis"`
		Sync Cache   `yaml:"sync"`
	} `yaml:"cache"`
}

type Cache struct {
	Host string		`yaml:"host"`
	Port string		`yaml:"port"`
	Pwd string	`yaml:"pwd"`
}
//其他包调用全局指针变量，如果没有初始化直接调用会报错
var Config *Configs

func Init() {
	//注意文件的相对位置
	//将来还可以通过请求consule获取配置文件
	content, err := ioutil.ReadFile("./config/app.yaml")
	if err != nil {
		panic(err)
	}
    Config = &Configs{}
	_ = yaml.Unmarshal(content, Config)
}