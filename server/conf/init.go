// Copyright 2016, Tonyxu All rights reserved.
// Author TonyXu <tonycbcd@gmail.com>
// Build on dev-0.0.1
// MIT Licensed

// To initialize the config data.

package conf

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	ConfigFile = flag.String("c", "./config/config.yml", "配置文件")
	InitConfig()
}

func InitConfig() {
	flag.Parse()
	if file, err := os.OpenFile(*ConfigFile, os.O_RDONLY, 0444); err != nil {
		log.Fatalln("配置文件读取失败", err)
	} else if in, err := ioutil.ReadAll(file); err != nil {
		log.Fatalln("配置文件读取失败", err)
	} else if err := yaml.Unmarshal(in, Config); err != nil {
		log.Fatalln("配置文件解析失败", err)
	}

}
