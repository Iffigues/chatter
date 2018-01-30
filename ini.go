package main

import(
	"log"
	"github.com/go-ini/ini"
)

const(
	bdd = "bdd"
	htp = "htp"
)

var(
	cfg, cfgErr = ini.Load("ini.ini")
)

func getKey(section , key string)(inu string){
	ar,err:=  cfg.Section(section).GetKey(key)
	if err != nil {
		log.Panic(err)
	}
	return ar.String()
}
