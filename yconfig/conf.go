package yconfig

import "github.com/BurntSushi/toml"

type Conf struct {
	CodeOutFile string
	Package     string
	Clickhouse  struct {
		Host string
		DB   string
		User string
		Pwd  string
		Port uint
		Imp []string
		Ext string
	}
}

var Config Conf

func ReadConf() error {
	_, err := toml.DecodeFile("./conf.toml", &Config)
	if err != nil {
		return err
	}
	return nil
}