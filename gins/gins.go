package gins

var Instance *Server

func init() {

	var err error
	Instance, err = New()
	if err != nil {
		panic("初始化服务失败!")
	}
}


func Run(conf *Config) {
	
	Instance.Init(conf)
    Instance.Start()
}