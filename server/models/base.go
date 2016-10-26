// Copyright 2016, Tonyxu All rights reserved.
// Author TonyXu <tonycbcd@gmail.com>
// Build on dev-0.0.1
// MIT Licensed

// The models.

package models

import (
    "fmt"
    "log"
    "time"
    "github.com/tonycbcd/easycms/server/conf"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
    DB *gorm.DB
)

type BaseModel struct {
    ID          uint `gorm:"primary_key"`
    Created     time.Time
    Updated     time.Time
    Deleted     *time.Time
    State       int         // 数据状态
}

func Init() {
    mysqlConf   := conf.Config.Mysql
    var err error
    DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", mysqlConf.User, mysqlConf.Pass, mysqlConf.Host, mysqlConf.Port, mysqlConf.Db))
    if err != nil {
        log.Fatalln("The DB connection error", err)
    }
    //DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
}

