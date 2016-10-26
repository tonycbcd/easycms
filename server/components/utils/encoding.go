// Copyright 2016, BeiJingWeSai Teq Inc. All rights reserved.
// Author Tonyxu <tonycbcd@gmail.com>,
// Build on dev-0.0.1
// MIT Licensed

package utils

import (
    "fmt"
    "crypto/md5"
)

type encoding struct{
}

var (
    PublicSalt  string
    oneEncoding *encoding
)

func NewEncoding() *encoding {
    if oneEncoding == nil {
        oneEncoding = &encoding{}
    }

    return oneEncoding
}

func (this *encoding) Md5(str string) string {
    return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
