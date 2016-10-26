// Copyright 2016, BeiJingWeSai Teq Inc. All rights reserved.
// Author Tonyxu <tonycbcd@gmail.com>,
// Build on dev-0.0.1
// MIT Licensed

package utils

import (
    "fmt"
    "errors"
    "strings"
    "strconv"
    "net/url"
    "net/http"
    "io/ioutil"
    "io"
    "bytes"
)

type SelfHttp struct{
}

type AuthPair struct {
    // user format is : AppId-Timestamp
    User    string

    // password format is : md5(Timstamp + AppPublickey)
    Pass    string
}

var (
    oneHttp *SelfHttp
)

func NewHttp() *SelfHttp {
    if oneHttp == nil {
        oneHttp = &SelfHttp{}
    }

    return oneHttp
}

func (this *SelfHttp) CloneParam(old *map[string]interface{}) *map[string]interface{} {
    newParams   := map[string]interface{}{}
    oldParams   := *old
    for k, v := range oldParams {
        newParams[ k ] = v
    }

    return &newParams
}

func (this *SelfHttp) GetLevelParamInRequest(allParam *map[string]interface{}) (*map[string]interface{}, error) {
    result := &map[string]interface{}{}
    var err error
    var res interface{}
    for k, v := range *allParam {
        findEd := []string{}
        for _, k1 := range strings.Split(k, "[") {
            if k1 != "" {
                for _, k2 := range strings.Split(k1, "]") {
                    if k2 != "" {
                        findEd = append(findEd, k2)
                    }
                }
            }
        }
        res, err = this.setResult(res, findEd, v)
        if err != nil {
            return result, err
        }
    }
    if paramsMap, ok := res.(map[string]interface{}); ok {
        result = &paramsMap
    }
    return result, nil
}


// to notice the master node change by http.
func (this *SelfHttp) NoticeMaster(addr string, mesg []byte) error {
    v := url.Values{}
    v.Set("node", string(mesg))

    body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
    client := &http.Client{}
    req, _ := http.NewRequest("POST", fmt.Sprintf("http://%s/notice/easyapinode", addr), body)
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
    resp, err := client.Do(req)
    if err == nil && resp != nil {
        defer resp.Body.Close()
    }

    return err
}

// to sync the slave node visited data to Master
func (this *SelfHttp) NoticeVisitedToMaster(addr string, mesg []byte) error {
    v := url.Values{}
    v.Set("part", string(mesg))

    body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
    client := &http.Client{}
    req, _ := http.NewRequest("POST", fmt.Sprintf("http://%s/notice/easyapivisited", addr), body)
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
    resp, err := client.Do(req)
    if err == nil && resp != nil {
        defer resp.Body.Close()
    }

    return err
}

// to ping the aim node.
func (this *SelfHttp) Ping(addr string) (resp string) {
    result  := ""
    defer func () {
        if err1 := recover(); err1 != nil {
            fmt.Printf("Ping err: %s\n", err1)
        }
    }()
    client := &http.Client{}
    if req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/notice/pingeasyapi", addr), nil); err != nil {
        return ""
    } else {
        req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
        resp, err := client.Do(req)
        if err == nil && resp != nil {
            defer resp.Body.Close()
            data, _ := ioutil.ReadAll(resp.Body)
            return string(data)
        }
    }

    return result
}

// to find request type.
func (this *SelfHttp) GetRequestAcceptType(data []string) string {
    dataStr := strings.ToLower(strings.Join(data, ":"))
    result  := "json"
    if strings.HasPrefix(dataStr, "application/xml") {
        result  = "xml"
    }
    return result
}

func (this *SelfHttp) InitRequest(body io.Reader) *http.Request {
    request := new(http.Request)
    rc, ok := body.(io.ReadCloser)
    if !ok && body != nil {
        rc = ioutil.NopCloser(body)
    }
    request.Body = rc
    request.Header = make(http.Header)
    if body != nil {
        switch v := body.(type) {
        case *bytes.Buffer:
            request.ContentLength = int64(v.Len())
        case *bytes.Reader:
            request.ContentLength = int64(v.Len())
        case *strings.Reader:
            request.ContentLength = int64(v.Len())
        }
    }
    return request
}

// 合并复合入参解析后数据.
func (this *SelfHttp) mergeParam(to, from interface{}) (interface{}, error) {
    switch to.(type) {
        case map[string]interface{}:
            if fromV, ok := from.(map[string]interface{}); ok {
                toV := to.(map[string]interface{})
                for k, v := range fromV {
                    if _, ok := toV[ k ]; ok {
                        return nil, errors.New(fmt.Sprintf("the param key '%s' was duplicate.", k))
                    } else {
                        toV[k]  = v
                    }
                }
                to  = toV
            } else {
                return nil, errors.New("merge param type does not match")
            }
        case []interface{}:
            if fromV, ok := from.([]interface{}); ok {
                toV := to.([]interface{})
                for _, v := range fromV {
                    toV = append(toV, v)
                }
                to      = toV
            } else {
                return nil, errors.New("merge param type does not match")
            }
    }

    return to, nil
}

func (this *SelfHttp) setResult (res interface{}, keys []string, value interface{}) (interface{}, error) {
    var err error
    if len(keys) == 0 {
        return value, nil
    }
    key := keys[0]
    if key != "" {
        isArr   := false
        arrIndex        := 0
        if arrIndex, err = strconv.Atoi(key); err == nil {
            isArr   = true
        }
        if res == nil {
            if isArr {
                res     = make([]interface{}, arrIndex + 1)
            } else {
                res     = map[string]interface{}{}
            }
        }

        switch res.(type) {
            case map[string]interface{}:
                tmp     := res.(map[string]interface{})
                if m, ok := tmp[key]; ok {
                    tmp[key], err = this.setResult(m, keys[1:], value)
                } else {
                    tmp[key], err = this.setResult(nil, keys[1:], value)
                }
                res     = tmp
            case []interface{}:
                tmp := res.([]interface{})
                var v interface{}
                if len(tmp) > arrIndex {
                    v = tmp[arrIndex]
                } else {
                    for i := 0; i <= arrIndex - len(tmp); i++ {
                        tmp = append(tmp, nil)
                    }
                }
                result, err := this.setResult(v, keys[1:], value)
                if err != nil {
                    return res, err
                }
                tmp[arrIndex] = result
                res = tmp
            default:
                err = errors.New("Parameter error")
                return nil, err
        }
    } else {
        if res == nil {
            res = []interface{}{}
        }
        tmp, ok := res.([]interface{})
        if !ok {
            err = errors.New("Parameter error")
            return res, err
        }
        var vv interface{}
        vv, err = this.setResult(res, keys[1:], value)
        if err != nil {
            return nil, err
        }
        tmp = append(tmp, vv)
        res = tmp
    }
    return res, nil
}
