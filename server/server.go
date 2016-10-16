// Copyright 2016, Tonyxu All rights reserved.
// Author TonyXu <tonycbcd@gmail.com>
// Build on dev-0.0.1
// MIT Licensed

// To initialize the config data.

package server

import (
    "github.com/tonycbcd/easycms/server/conf"
    "fmt"
)

func Run() {
    fmt.Printf("config: %#v\n", conf.Config)
}
