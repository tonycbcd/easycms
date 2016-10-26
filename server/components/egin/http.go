package egin

import (
    "strings"
    "encoding/base64"
    "errors"
)

type Http struct {
    c   *Context
}

type AuthPair struct {
    // user format is : AppId-Timestamp
    User    string
    // password format is : md5(Timstamp + AppPublickey)
    Pass    string
}

func (this *Http) ParseBasicAuthPair() (*AuthPair, error) {
    basicAuthor := this.c.Request.Header.Get("Authorization")
    basicAuthor = strings.Replace(basicAuthor, "Basic ", "", 1)
    if authorStr, err := base64.StdEncoding.DecodeString(basicAuthor); err != nil {
        return nil, err
    } else if pair := strings.Split(string(authorStr), ":"); len(pair) > 1 {
        return &AuthPair{User: pair[0], Pass: pair[1]}, nil
    } else {
        return nil, errors.New("no any author")
    }
}

func (this *Http) ClientIP() string {
    ipInfo := strings.Split(this.c.ClientIP(), ":")
    return ipInfo[0]
}
