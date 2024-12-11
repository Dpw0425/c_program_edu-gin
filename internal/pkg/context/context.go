package ctx

import (
	"c_program_edu-gin/internal/app/middleware"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"runtime"
)

// MarshalOptions is a configurable JSON format marshaller.
var MarshalOptions = protojson.MarshalOptions{
	UseProtoNames:   true,
	EmitUnpopulated: true,
}

type Context struct {
	Context *gin.Context
}

func New(ctx *gin.Context) *Context {
	return &Context{Context: ctx}
}

// JSONData 将 proto 对象转为 json 对象
func (c *Context) JSONData(data any) interface{} {
	if value, ok := data.(proto.Message); ok {
		b, _ := MarshalOptions.Marshal(value)

		var body map[string]any
		_ = json.Unmarshal(b, &body)

		return body
	}

	return nil
}

// UserID 返回登录用户的 uid
func (c *Context) UserID() int {
	if session := c.JWTSession(); session != nil {
		return session.Uid
	}

	return 0
}

// JWTSession 返回登录用户的 jwt session
func (c *Context) JWTSession() *middleware.JSession {
	tokenString, ok := c.Context.Get(middleware.JWTSessionConst)
	if !ok {
		return nil
	}

	return tokenString.(*middleware.JSession)
}

// IsGuest 判断是否登录 (游客状态)
func (c *Context) IsGuest() bool {
	return c.UserID() == 0
}

func (c *Context) Ctx() context.Context {
	return c.Context.Request.Context()
}

// 返回错误出现的行号
func initMeta() map[string]any {
	meta := make(map[string]any)

	_, _, line, ok := runtime.Caller(2)
	if ok {
		meta["error_line"] = line
	}

	return meta
}
