package middleware

import (
	"bytes"
	"c_program_edu-gin/utils/json_utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"slices"
	"strings"
	"time"
)

type AccessLogStore struct {
	ctx       *gin.Context
	startTime time.Time
	req       *RequestInfo
}

func newAccessLogStore(c *gin.Context) *AccessLogStore {
	return &AccessLogStore{
		ctx:       c,
		startTime: time.Now(),
		req:       nil,
	}
}

// 初始化过滤器日志
func (a *AccessLogStore) init() error {
	hostname, _ := os.Hostname()

	headers := make(map[string]string)
	for k := range a.ctx.Request.Header {
		headers[k] = a.ctx.Request.Header.Get(k)
	}

	body, err := io.ReadAll(a.ctx.Request.Body)
	if err != nil {
		return err
	}

	a.ctx.Request.Body = io.NopCloser(bytes.NewReader(body))

	a.req = &RequestInfo{
		HostName:        hostname,
		RemoteAddr:      a.ctx.RemoteIP(),
		ServerName:      a.ctx.Request.Host,
		HttpUserAgent:   a.ctx.Request.UserAgent(),
		RequestID:       a.ctx.Request.Header.Get("X-Request-ID"),
		RequestTime:     a.startTime.Format("2024-05-03 19:09:45"),
		RequestMethod:   a.ctx.Request.Method,
		RequestHeader:   headers,
		RequestURI:      a.ctx.Request.URL.Path,
		RequestQuery:    urlValuesToMap(a.ctx.Request.URL.Query()),
		RequestBody:     string(body),
		ResponseHeader:  make(map[string]string),
		ResponseBody:    "",
		RequestBodyRaw:  "",
		ResponseTime:    "",
		HttpStatus:      0,
		RequestDuration: "",
		MetaData:        make(map[string]any),
	}

	if a.req.RequestID == "" {
		a.req.RequestID = uuid.New().String()
	}

	return nil
}

// 日志内容配置
func (a *AccessLogStore) load() {
	writer := a.ctx.Writer.(responseWriter)

	headers := make(map[string]string)
	for k := range writer.Header() {
		headers[k] = writer.Header().Get(k)
	}

	a.req.ResponseHeader = headers
	a.req.ResponseTime = time.Now().Format("2024-05-03 19:28:29")
	a.req.RequestDuration = fmt.Sprintf("%.3f", time.Since(a.startTime).Seconds())
	a.req.HttpStatus = writer.Status()
	a.req.ResponseBody = writer.body.String()
	a.req.ResponseBodyRaw = a.req.ResponseBody

	// SESSION JWT
	session, ok := a.ctx.Get(JWTSessionConst)
	if ok {
		a.req.MetaData["uid"] = session.(*JSession).Uid
	}
}

// 保存日志
func (a *AccessLogStore) save(log *slog.Logger) {
	data := make(map[string]any)
	if err := json_utils.Decode(json_utils.Encode(a.req), &data); err != nil {
		return
	}

	// 解析 request
	if strings.HasPrefix(a.ctx.GetHeader("Content-Type"), "application/json_utils") {
		var body map[string]any
		_ = json.Unmarshal([]byte(a.req.RequestBody), &body)

		data["request_body"] = body
		delete(data, "request_body_raw")
	} else {
		delete(data, "request_body")
	}

	// 解析 response
	writer := a.ctx.Writer.(responseWriter)
	if strings.HasPrefix(writer.Header().Get("Content-Type"), "application/json") {
		var body map[string]any
		_ = json.Unmarshal([]byte(a.req.ResponseBody), &body)

		data["response_body"] = body
		delete(data, "response_body_raw")
	} else {
		delete(data, "response_body")
	}

	items := make([]any, 0)
	for k, v := range data {
		items = append(items, k, v)
	}

	log.With(items...).Info("access_log")
}

func AccessLog(w io.Writer, filterRule *AccessFilterRule) gin.HandlerFunc {
	// 初始化日志配置
	log := slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(a.Value.Time().Format("2006-01-02 15:04:05.000"))
			}

			return a
		},
	}))

	return func(c *gin.Context) {
		// 捕获响应数据并初始化
		c.Writer = responseWriter{c.Writer, bytes.NewBuffer([]byte{})}

		// 记录当前请求的访问日志
		access := newAccessLogStore(c)
		if err := access.init(); err != nil {
			c.Abort()
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
			return
		}

		c.Next()

		// 处理请求后写入日志
		if c.Request.Method != "OPTIONS" {
			access.load()

			if filterRule != nil {
				filterRule.filter(access)
			}

			if slices.Contains(filterRule.excludeRoutes, access.req.RequestURI) {
				return
			}

			access.save(log)
		}
	}
}

func urlValuesToMap(values url.Values) map[string]any {
	data := make(map[string]any)
	for k, v := range values {
		if len(v) == 1 {
			data[k] = v[0]
		} else {
			data[k] = v
		}
	}
	return data
}
