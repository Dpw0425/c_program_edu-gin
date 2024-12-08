package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

type RequestInfo struct {
	HostName        string            `json_utils:"host_name"`
	RemoteAddr      string            `json_utils:"remote_addr"`
	ServerName      string            `json_utils:"server_name"`
	HttpUserAgent   string            `json_utils:"http_user_agent"`
	RequestID       string            `json_utils:"request_id"`
	RequestTime     string            `json_utils:"request_time"`
	RequestMethod   string            `json_utils:"request_method"`
	RequestHeader   map[string]string `json_utils:"request_header"`
	RequestURI      string            `json_utils:"request_uri"`
	RequestQuery    map[string]any    `json_utils:"request_query"`
	RequestBody     string            `json_utils:"request_body"`
	RequestBodyRaw  string            `json_utils:"request_body_raw"`
	ResponseHeader  map[string]string `json_utils:"response_header"`
	ResponseBody    string            `json_utils:"response_body"`
	ResponseBodyRaw string            `json_utils:"response_body_raw"`
	ResponseTime    string            `json_utils:"response_time"`
	HttpStatus      int               `json_utils:"http_status"`
	RequestDuration string            `json_utils:"request_duration"`
	MetaData        map[string]any    `json_utils:"meta_data"`
}

// 过滤行为实体
type AccessFilterOption struct {
	path string
	fn   func(req *RequestInfo)
}

// 过滤规则
type AccessFilterRule struct {
	opts          map[string]AccessFilterOption
	excludeRoutes []string
}

// 新建过滤规则实例
func NewAccessFilterRule() *AccessFilterRule {
	return &AccessFilterRule{opts: make(map[string]AccessFilterOption), excludeRoutes: make([]string, 0)}
}

func (rule *AccessFilterRule) Exclude(path string) {
	rule.excludeRoutes = append(rule.excludeRoutes, path)
}

func (rule *AccessFilterRule) AddRule(path string, fn func(req *RequestInfo)) {
	rule.opts[path] = AccessFilterOption{
		path: path,
		fn:   fn,
	}
}

func (rule *AccessFilterRule) filter(access *AccessLogStore) {
	if opt, ok := rule.opts[access.req.RequestURI]; ok {
		opt.fn(access.req)
	}
}

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
