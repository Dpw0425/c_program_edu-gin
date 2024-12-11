package ctx

import (
	"c_program_edu-gin/pkg/response"
	"github.com/gin-gonic/gin"
)

func HandlerFunc(fn func(ctx *Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := fn(New(c)); err != nil {
			response.ErrResponse(c, err)
			return
		}
	}
}
