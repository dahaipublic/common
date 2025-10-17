package blademaster

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Param map[string]interface{}

type Header map[string]string

type Req struct {
	Url string
	Param
	Header
}

type result struct {
	Url string
	Param
	Method     string
	Params     interface{}
	Body       string
	Header     http.Header
	Response   *http.Response
	Proto      string
	Host       string
	URL        *url.URL
	Status     string
	StatusCode int
}

// 模拟请求
func (r *Req) do(method string) (ret *result, err error) {
	reqs, err := http.NewRequest(method, string(r.Url), nil)
	if err != nil {
		return nil, err
	}
	// 添加header
	r.setHeader(reqs, r.Header)
	// 设置参数
	q := reqs.URL.Query()
	for k, v := range r.Param {
		q.Add(k, fmt.Sprint(v))
	}
	reqs.URL.RawQuery = q.Encode()
	res, err := http.DefaultClient.Do(reqs)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	ret = &result{
		Url:        r.Url,
		Param:      r.Param,
		Method:     method,
		Body:       r.getBody(res),
		Header:     reqs.Header,
		Response:   res,
		Proto:      reqs.Proto,
		Host:       reqs.Host,
		URL:        reqs.URL,
		Status:     res.Status,
		StatusCode: res.StatusCode,
	}
	return
}

func (r *Req) postJson() (ret *result, err error) {
	jsonStr, _ := json.Marshal(r.Param)
	req, err := http.NewRequest("POST", r.Url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	// 添加header
	r.setHeader(req, r.Header)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ret = &result{
		Url:        r.Url,
		Param:      r.Param,
		Method:     "POST",
		Body:       r.getBody(resp),
		Header:     req.Header,
		Response:   resp,
		Proto:      req.Proto,
		Host:       req.Host,
		URL:        req.URL,
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
	}
	return
}

// 获取body内容
func (r *Req) getBody(res *http.Response) (body string) {
	if res.StatusCode == 200 {
		switch res.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(res.Body)
			for {
				buf := make([]byte, 1024)
				n, err := reader.Read(buf)
				if err != nil && err != io.EOF {
					panic(err)
				}
				if n == 0 {
					break
				}
				body += string(buf)
			}
		default:
			bodyByte, _ := ioutil.ReadAll(res.Body)
			body = string(bodyByte)
		}
	} else {
		bodyByte, _ := ioutil.ReadAll(res.Body)
		body = string(bodyByte)
	}
	return
}

// 设置header
func (r *Req) setHeader(h *http.Request, header Header) {
	if len(header) > 0 {
		for key, value := range header {
			h.Header.Add(key, value)
		}
	}
	h.Header.Add("Content-Type", "text/plain; charset=UTF-8")
	h.Header.Add("User-Agent", "Go-http-client/1.14")
	h.Header.Add("Transfer-Encoding", "chunked")
	h.Header.Add("Accept-Encoding", "gzip, deflate")
}

//设置参数

// 发起请求
func (r *Req) Request(method string, response interface{}) (err error) {
	var ret *result
	switch method {
	case http.MethodGet:
		ret, err = r.do(http.MethodGet)
	case http.MethodPost:
		ret, err = r.do(http.MethodPost)
	case "POSTJSON":
		ret, err = r.postJson()
	case http.MethodDelete:
		ret, err = r.do(http.MethodDelete)
	case http.MethodPut:
		ret, err = r.do(http.MethodPut)
	}
	if err != nil {
		err = errors.New("请求失败")
		return
	}
	if ret.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("statusCode:%d", ret.StatusCode))
		return
	}

	err = json.Unmarshal([]byte(ret.Body), &response)
	if err != nil {
		return
	}
	return
}
