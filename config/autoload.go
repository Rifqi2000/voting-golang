package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type CServer struct {
	Server struct {
		Address string `json:"address"`
		Port    string `json:"port"`
	} `json:"server"`

	Db struct {
		Driver   string `json:"driver"`
		Address  string `json:"address"`
		Port     string `json:"port"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
		Prefix   string `json:"prefix"`
	} `json:"db"`
}

var c_server *CServer

var dirname string = "config"
var flname_cserver string = "server.json"

func init() {
	if c_server != nil {
		return
	}

	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
		return
	}

	bts, err := ioutil.ReadFile(filepath.Join(basePath, dirname, flname_cserver))
	if err != nil {
		panic(err)
		return
	}

	c_server = new(CServer)
	err = json.Unmarshal(bts, &c_server)
	if err != nil {
		panic(err)
		return
	}
}

func GetCServer() CServer {
	return *c_server
}
