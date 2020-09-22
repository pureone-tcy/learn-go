package main

import (
	"fmt"
	"github.com/go-ini/ini"
)

// ini Config
type ConfigList struct {
	Port      int
	DBName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("./ex7/config.ini")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustInt(),
		DBName:    cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("%T %v \n", Config.Port, Config.Port)
	fmt.Printf("%T %v \n", Config.DBName, Config.DBName)
	fmt.Printf("%T %v \n", Config.SQLDriver, Config.SQLDriver)
}
