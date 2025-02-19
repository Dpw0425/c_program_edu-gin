package middleware

import (
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/pkg/jwt"
	"c_program_edu-gin/pkg/response"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

const JWTSessionConst = "__JWT_SESSION__"

type JSession struct {
	Uid       int64  `json_utils:"uid"`
	Token     string `json_utils:"token"`
	ExpiresAt int64  `json_utils:"expires_at"`
}

var (
	ErrorNoLogin = myErr.Unauthorized("", "未登录！")
)

type IStorage interface {
	// IsBlackList 判断是否是黑名单
	IsBlackList(ctx context.Context, token string) bool
}

// 鉴权中间件
func Auth(secret string, guard string, storage IStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := GetHeaderToken(c)

		claims, err := verify(guard, secret, token)
		if err != nil {
			response.ErrResponse(c, err)
			return
		}

		if storage.IsBlackList(c.Request.Context(), token) {
			response.ErrResponse(c, myErr.Unauthorized("", "登录已过期！"))
		}

		uid, err1 := strconv.ParseInt(claims.ID, 10, 64)
		if err1 != nil {
			response.ErrResponse(c, myErr.InternalServerError("", "解析 JWT 失败！"))
			return
		}

		c.Set(JWTSessionConst, &JSession{
			Uid:       uid,
			Token:     token,
			ExpiresAt: claims.ExpiresAt.Unix(),
		})

		c.Next()
	}
}

// 获取 token
func GetHeaderToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer"))

	if token == "" {
		token = c.DefaultQuery("token", "")
	}

	return token
}

func verify(guard string, secret string, token string) (*jwt.AuthClaims, error) {
	if token == "" {
		return nil, ErrorNoLogin
	}

	claims, err := jwt.ParseToken(token, secret)
	if err != nil {
		return nil, err
	}

	if claims.Guard != guard || claims.Valid() != nil {
		return nil, ErrorNoLogin
	}

	return claims, nil
}
