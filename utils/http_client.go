package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	netUrl "net/url"
	"strings"
	"time"
)

type HttpOption func(*httpRequest)

type httpRequest struct {
	client  *http.Client
	url     string
	method  string
	headers map[string]string
	timeout time.Duration
}

// 设置超时时间选项
func WithTimeout(timeout time.Duration) HttpOption {
	return func(r *httpRequest) {
		r.timeout = timeout
	}
}

// 设置请求头选项
func WithHeaders(headers map[string]string) HttpOption {
	return func(r *httpRequest) {
		r.headers = headers
	}
}

// 发送HTTP请求
func HttpDo(url string, method string, params map[string]interface{}, opts ...HttpOption) ([]byte, error) {
	// 默认配置
	req := &httpRequest{
		client:  &http.Client{},
		url:     url,
		method:  method,
		headers: make(map[string]string),
		timeout: 30 * time.Second,
	}

	// 应用选项
	for _, opt := range opts {
		opt(req)
	}

	// 设置超时
	req.client.Timeout = req.timeout

	var body *bytes.Buffer
	// 处理请求参数
	if params != nil {
		if method == http.MethodGet {
			// GET请求：将参数添加到URL查询字符串
			query := req.url
			if len(params) > 0 {
				if !strings.Contains(query, "?") {
					query += "?"
				} else if !strings.HasSuffix(query, "&") {
					query += "&"
				}
				values := netUrl.Values{}
				for k, v := range params {
					values.Add(k, fmt.Sprintf("%v", v))
				}
				query += values.Encode()
			}
			req.url = query
			body = bytes.NewBuffer(nil)
		} else {
			// POST请求：将参数转换为JSON
			jsonData, err := json.Marshal(params)
			if err != nil {
				return nil, err
			}
			body = bytes.NewBuffer(jsonData)
			// 设置默认的Content-Type
			if _, exists := req.headers["Content-Type"]; !exists {
				req.headers["Content-Type"] = "application/json"
			}
		}
	}
	if body == nil {
		body = bytes.NewBuffer(nil)
	}
	// 创建请求
	httpReq, err := http.NewRequest(req.method, req.url, body)
	if err != nil {
		return nil, err
	}

	// 设置请求头
	for k, v := range req.headers {
		httpReq.Header.Set(k, v)
	}

	// 发送请求
	response, err := req.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(response.Body)
}
