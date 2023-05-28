package main

import (
	"apischool/iternal/app/apiserver"
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configpath string
)

func init(){
	flag.StringVar(&configpath, "config-path", "iternal/app/apiserver/configs/apisever.toml", "path to config file")

}
func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configpath, config)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.NEw(config)
	if err := s.Start(); err != nil{
		log.Fatal(err)
	}
}

