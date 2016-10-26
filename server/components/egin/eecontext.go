package egin

import (
    "github.com/gin-gonic/gin"
    "bytes"
    "strconv"
    "strings"
    "crypto/md5"
    "fmt"
)

var (
    JsonInput bool = false
)

type Context struct {
    *gin.Context
    BodyBuffer      *bytes.Buffer
    UserAuthKey     string

    ContentType     string
    RequestJson     string

    // 请求头信息中的版本(做兼容用)
    Version         int64
    httpOne         *Http
}

type H gin.H

func NewContext(c *gin.Context) *Context {
    var e *Context
    if i, ok := c.Get("__EGIN_CONTEXT__"); ok {
        e = i.(*Context)
    } else {
        e = &Context{
            Context: c,
        }
        e.setVersion()
        contentType := c.Request.Header.Get("Content-Type")
        // 如果版本好号大于等于2.0.11 才支持json入参
        if contentType != "" && strings.Index(strings.ToLower(contentType), "application/json") != -1 && (e.Version >= 2000000011 || JsonInput) {
            e.ContentType = "JSON"
        } else if strings.Index(strings.ToLower(contentType), "text/xml") != -1 {
            e.ContentType = "XML"
        }
        c.Set("__EGIN_CONTEXT__", e)
    }
    return e
}

func (this *Context) setVersion() {
    version := this.Request.Header.Get("Apiversion")
    if version != "" {
        tmp := strings.Split(version, ".")
        var (
            v1, v2, v3 int64
        )
        if len(tmp) > 0 {
            v1, _ = strconv.ParseInt(tmp[0], 10, 64)
        }
        if len(tmp) > 1 {
            v2, _ = strconv.ParseInt(tmp[1], 10, 64)
        }
        if len(tmp) > 2 {
            v3, _ = strconv.ParseInt(tmp[2], 10, 64)
        }
        this.Version = v1 * 1000000000 + v2 * 1000000 + v3
    }
}

func (this *Context) SaveUserAuthKey(userAuthKey string) {
    this.UserAuthKey = userAuthKey
}

func (this *Context) NewHttp() *Http {
    if this.httpOne == nil {
        this.httpOne = &Http{c:this}
    }
    return this.httpOne
}

func Md5(str string) string {
    return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
