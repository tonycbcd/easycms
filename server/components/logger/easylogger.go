// Copyright 2015, WePiao.com Teq Inc. All rights reserved.
// Build on dev-0.0.1
// MIT Licensed

package logger

import (
    "fmt"
    "log"
    "os"
)

const (
    VERSION     = "2.0.16"
)

var (
	Shell       = log.New(os.Stdout, fmt.Sprintf("[EASY API(v%s)]:", VERSION), log.LstdFlags | log.Lshortfile)

    Proxy       = Shell
)


