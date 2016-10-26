// Copyright 2016, BeiJingWeSai Teq Inc. All rights reserved.
// Author Tonyxu <tonycbcd@gmail.com>,
// Build on dev-0.0.1
// MIT Licensed

package utils

import (
    "testing"
    "fmt"
)

func Test_http(t *testing.T) {
    allParam := map[string]interface{}{
        "a[name]": 100,
        "a[d][0][certNo]": 1,
        "a[d][0][certName]": "李四",
        "a[d][1][certNo]": 2,
        "a[d][1][certName]": "张三",
        "a[d][2][certName]": "张三",
    }
    v, err := NewHttp().GetLevelParamInRequest(&allParam)
    fmt.Println("结果", v)
    fmt.Println("错误", err)
}
